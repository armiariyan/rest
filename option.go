package rest

import (
	"errors"
)

// Option configures Client with defined option.
type Option func(*DefaultHttpRequester) error

// WithCircuitBreaker returns Option to configure Circuit Breaker
func WithCircuitBreaker(circuitConfig CBConfig) Option {
	return func(c *DefaultHttpRequester) error {
		c.client = newCircuitBreaker(circuitConfig, c.client)
		return nil
	}
}

// AddHook returns Option to adding new hook
func AddHook(hook Hook) Option {
	return func(c *DefaultHttpRequester) error {
		if hook == nil {
			return errors.New("hook is nil")
		}

		c.hook = append(c.hook, hook)
		return nil
	}
}
