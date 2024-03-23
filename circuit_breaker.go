package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sony/gobreaker"
)

// State is a type that represents a state of CircuitBreaker.
type State int

// These constants are states of CircuitBreaker.
const (
	StateClosed State = iota
	StateHalfOpen
	StateOpen
)

type ReadyToTripFunc func(Counts) bool
type OnStateChangeFunc func(name string, from State, to State)

// CBConfig configures CircuitBreaker:
//
// Name is the name of the CircuitBreaker.
//
// MaxRequests is the maximum number of requests allowed to pass through
// when the CircuitBreaker is half-open.
// If MaxRequests is 0, the CircuitBreaker allows only 1 request.
//
// Interval is the cyclic period of the closed state
// for the CircuitBreaker to clear the internal Counts.
// If Interval is 0, the CircuitBreaker doesn't clear internal Counts during the closed state.
//
// Timeout is the period of the open state,
// after which the state of the CircuitBreaker becomes half-open.
// If Timeout is 0, the timeout value of the CircuitBreaker is set to 60 seconds.
//
// ReadyToTrip is called with a copy of Counts whenever a request fails in the closed state.
// If ReadyToTrip returns true, the CircuitBreaker will be placed into the open state.
// If ReadyToTrip is nil, default ReadyToTrip is used.
// Default ReadyToTrip returns true when the number of consecutive failures is more than 5.
//
// OnStateChange is called whenever the state of the CircuitBreaker changes.
type CBConfig struct {
	Name            string
	IsActive        bool
	Timeout         int
	IntervalTimeout int
	Threshold       int
	Paths           []string
	MaxRequests     uint32
	ReadyToTrip     ReadyToTripFunc
	OnStateChange   OnStateChangeFunc
}

// Counts holds the numbers of requests and their successes/failures.
// CircuitBreaker clears the internal Counts either
// on the change of the state or at the closed-state intervals.
// Counts ignores the results of the requests sent before clearing.
// This copies from gobreaker.Counts so that user don't need to know the external lib.
type Counts struct {
	Requests             uint32
	TotalSuccesses       uint32
	TotalFailures        uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
}

type circuitBreaker struct {
	client         HttpClient
	useBreaker     bool
	breaker        *gobreaker.CircuitBreaker
	whitelistPaths []string
}

// readyToTrip wraps gobreaker.ReadyToTrip function
var readyToTrip = func(setting ReadyToTripFunc) func(gobreaker.Counts) bool {
	if setting == nil {
		return nil
	}

	return func(counts gobreaker.Counts) bool {
		return setting(Counts{
			Requests:             counts.Requests,
			TotalSuccesses:       counts.TotalSuccesses,
			TotalFailures:        counts.TotalFailures,
			ConsecutiveSuccesses: counts.ConsecutiveSuccesses,
			ConsecutiveFailures:  counts.ConsecutiveFailures,
		})
	}
}

// onStateChange wraps gobreaker.State listener
var onStateChange = func(stateListener OnStateChangeFunc) func(name string, from gobreaker.State, to gobreaker.State) {
	if stateListener == nil {
		return nil
	}

	return func(name string, from gobreaker.State, to gobreaker.State) {
		var (
			internalFrom State
			internalTo   State
		)

		switch from {
		case gobreaker.StateClosed:
			internalFrom = StateClosed
		case gobreaker.StateHalfOpen:
			internalFrom = StateHalfOpen
		case gobreaker.StateOpen:
			internalFrom = StateOpen
		}

		switch to {
		case gobreaker.StateClosed:
			internalTo = StateClosed
		case gobreaker.StateHalfOpen:
			internalTo = StateHalfOpen
		case gobreaker.StateOpen:
			internalTo = StateOpen
		}

		stateListener(name, internalFrom, internalTo)
	}
}

func newCircuitBreaker(conf CBConfig, client HttpClient) *circuitBreaker {
	cb := new(circuitBreaker)
	cb.client = client
	cb.useBreaker = conf.IsActive
	cb.whitelistPaths = conf.Paths
	cb.breaker = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:          conf.Name,
		MaxRequests:   conf.MaxRequests,
		Interval:      time.Duration(conf.IntervalTimeout) * time.Second,
		Timeout:       time.Duration(conf.Timeout) * time.Second,
		ReadyToTrip:   readyToTrip(conf.ReadyToTrip),
		OnStateChange: onStateChange(conf.OnStateChange),
	})

	return cb
}

func (cb *circuitBreaker) Do(request *http.Request) (*http.Response, error) {
	if cb.useBreaker && cb.pathInWhitelist(request.URL.String()) {
		cbResp, err := cb.breaker.Execute(func() (interface{}, error) {
			resp, err := cb.client.Do(request) // resp should be nil when err not nil
			if err != nil {
				// https://github.com/golang/go/blob/2d77d3330537e11a0d9a233ba5f4facf262e9d8c/src/net/http/client.go#L724
				if strings.Contains(err.Error(), "Client.Timeout") {
					err = ErrHttpTimeout
					return resp, err
				}

				return resp, err
			}

			if resp.StatusCode >= http.StatusInternalServerError {
				return resp, fmt.Errorf("error request http status: %d", resp.StatusCode)
			}

			return resp, nil
		})

		var resp *http.Response
		if cbResp != nil {
			resp = cbResp.(*http.Response)
		}

		// https://github.com/golang/go/blob/2d77d3330537e11a0d9a233ba5f4facf262e9d8c/src/net/http/client.go#L724
		if err != nil && strings.Contains(err.Error(), "Client.Timeout") {
			err = ErrHttpTimeout
			return resp, err
		}

		return resp, err
	}

	return cb.client.Do(request)
}

func (cb *circuitBreaker) pathInWhitelist(path string) bool {
	for _, wp := range cb.whitelistPaths {
		u, _ := url.Parse(path)
		path = u.Path
		if wp == path {
			return true
		}
	}

	return false
}
