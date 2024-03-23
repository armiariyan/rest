package rest

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/sony/gobreaker"

	"github.com/smartystreets/goconvey/convey"
)

func TestReadyToTrip(t *testing.T) {
	convey.Convey("Ready to trip function", t, func() {
		convey.Convey("Should return nil, when callback func is nil", func() {
			rtt := readyToTrip(nil)
			convey.So(rtt, convey.ShouldBeNil)
		})

		convey.Convey("Should return true", func() {
			rtt := readyToTrip(func(counts Counts) bool {
				convey.So(counts.Requests, convey.ShouldResemble, uint32(1))
				convey.So(counts.TotalSuccesses, convey.ShouldResemble, uint32(2))
				convey.So(counts.TotalFailures, convey.ShouldResemble, uint32(3))
				convey.So(counts.ConsecutiveSuccesses, convey.ShouldResemble, uint32(4))
				convey.So(counts.ConsecutiveFailures, convey.ShouldResemble, uint32(5))
				return true
			})(gobreaker.Counts{
				Requests:             1,
				TotalSuccesses:       2,
				TotalFailures:        3,
				ConsecutiveSuccesses: 4,
				ConsecutiveFailures:  5,
			})

			convey.So(rtt, convey.ShouldBeTrue)
		})
	})
}

func TestOnStateChange(t *testing.T) {
	convey.Convey("onStateChange function", t, func() {
		convey.Convey("Should return nil, when callback func is nil", func() {
			osc := onStateChange(nil)
			convey.So(osc, convey.ShouldBeNil)
		})

		convey.Convey("Should return state closed", func() {
			osc := onStateChange(func(name string, from State, to State) {
				convey.So(name, convey.ShouldResemble, "a")
				convey.So(from, convey.ShouldResemble, StateClosed)
				convey.So(to, convey.ShouldResemble, StateClosed)
			})

			osc("a", gobreaker.StateClosed, gobreaker.StateClosed)
		})

		convey.Convey("Should return state half open", func() {
			osc := onStateChange(func(name string, from State, to State) {
				convey.So(name, convey.ShouldResemble, "a")
				convey.So(from, convey.ShouldResemble, StateHalfOpen)
				convey.So(to, convey.ShouldResemble, StateHalfOpen)
			})

			osc("a", gobreaker.StateHalfOpen, gobreaker.StateHalfOpen)
		})

		convey.Convey("Should return state open", func() {
			osc := onStateChange(func(name string, from State, to State) {
				convey.So(name, convey.ShouldResemble, "a")
				convey.So(from, convey.ShouldResemble, StateOpen)
				convey.So(to, convey.ShouldResemble, StateOpen)
			})

			osc("a", gobreaker.StateOpen, gobreaker.StateOpen)
		})
	})
}

func TestNewCircuitBreaker(t *testing.T) {
	convey.Convey("New Circuit Breaker", t, func() {
		convey.Convey("Should return not error", func() {
			testClient := &mockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					// do whatever you want
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			}

			cb := newCircuitBreaker(CBConfig{}, testClient)
			convey.So(cb, convey.ShouldNotBeNil)
		})
	})
}

func TestCircuitBreakerDo(t *testing.T) {
	convey.Convey("New Circuit Breaker", t, func() {
		request := &http.Request{}
		request.URL = &url.URL{
			Path: "/",
		}

		convey.Convey("Using circuit breaker, return success", func() {
			testClient := &mockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					// do whatever you want
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			}

			cb := newCircuitBreaker(CBConfig{
				IsActive: true,
				Paths:    []string{"/"},
			}, testClient)
			convey.So(cb, convey.ShouldNotBeNil)

			resp, err := cb.Do(request)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})

		convey.Convey("Using circuit breaker, return error", func() {
			testClient := &mockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					// do whatever you want
					return &http.Response{
						StatusCode: http.StatusOK,
					}, fmt.Errorf("error http")
				},
			}

			cb := newCircuitBreaker(CBConfig{
				IsActive: true,
				Paths:    []string{"/"},
			}, testClient)
			convey.So(cb, convey.ShouldNotBeNil)

			resp, err := cb.Do(request)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})

		convey.Convey("Using circuit breaker, return server 500", func() {
			testClient := &mockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					// do whatever you want
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
					}, nil
				},
			}

			cb := newCircuitBreaker(CBConfig{
				IsActive: true,
				Paths:    []string{"/"},
			}, testClient)
			convey.So(cb, convey.ShouldNotBeNil)

			resp, err := cb.Do(request)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})

		convey.Convey("Using circuit breaker, return Client.Timeout", func() {
			testClient := &mockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					var err = errors.New("error from nowhere")
					return nil, fmt.Errorf(err.Error() + " (Client.Timeout exceeded while awaiting headers)")
				},
			}

			cb := newCircuitBreaker(CBConfig{
				IsActive: true,
				Paths:    []string{"/"},
			}, testClient)
			convey.So(cb, convey.ShouldNotBeNil)

			resp, err := cb.Do(request)
			convey.So(resp, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldResemble, ErrHttpTimeout)
		})

	})
}

func TestCircuitBreakerDoWithoutCircuitBreaker(t *testing.T) {
	convey.Convey("Not using circuit breaker", t, func() {
		request := &http.Request{}
		request.URL = &url.URL{
			Path: "/",
		}

		convey.Convey("Should return success", func() {
			testClient := &mockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					// do whatever you want
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			}

			cb := newCircuitBreaker(CBConfig{}, testClient)
			convey.So(cb, convey.ShouldNotBeNil)

			resp, err := cb.Do(request)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestPathInWhitelist(t *testing.T) {
	convey.Convey("New Circuit Breaker", t, func() {
		convey.Convey("Should return true when path is in list", func() {

			cb := &circuitBreaker{
				whitelistPaths: []string{"/path"},
			}
			convey.So(cb, convey.ShouldNotBeNil)

			pathInWhiteList := cb.pathInWhitelist("/path")
			convey.So(pathInWhiteList, convey.ShouldBeTrue)
		})

		convey.Convey("Should return false when path is not in list", func() {

			cb := &circuitBreaker{
				whitelistPaths: []string{"/path"},
			}
			convey.So(cb, convey.ShouldNotBeNil)

			pathInWhiteList := cb.pathInWhitelist("/")
			convey.So(pathInWhiteList, convey.ShouldBeFalse)
		})
	})
}
