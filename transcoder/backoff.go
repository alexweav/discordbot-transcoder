package transcoder

import "time"

// The default interval for a constant back-off.
const DefaultConstantBackoffInterval = 10 * time.Second

// Provides a back-off policy for retrying some operation.
type Backoff interface {
	Next() time.Duration
	Reset()
}

// A back-off policy where retries occur at the same constant interval.
type constantBackoff struct {
	interval time.Duration
}

// Specifies the rate of a constant back-off policy.
type ConstantBackoffPolicy struct {
	Interval time.Duration
}

// Builds a new constant back-off policy, where retries occur at the same interval.
func NewConstantBackoff(policy ConstantBackoffPolicy) Backoff {
	interval := policy.Interval
	if interval < 1 {
		interval = DefaultConstantBackoffInterval
	}
	return &constantBackoff{
		interval: interval,
	}
}

// Gets the next back-off interval from a constant backoff.
func (backoff *constantBackoff) Next() time.Duration {
	return backoff.interval
}

// Resets a constant backoff.
func (backoff *constantBackoff) Reset() {}

/*
A back-off policy which increases the retry interval for each retry attempt
using a randomized function which grows exponentially.*/
type exponentialBackoff struct {
	currentInterval     time.Duration
	initialInterval     time.Duration
	randomizationFactor float64
	multiplier          float64
	maxInterval         time.Duration
	maxElapsedTime      time.Duration
}

// Specifies the rate of an exponential back-off policy.
type ExponentialBackoffPolicy struct {
	InitialInterval     time.Duration
	RandomizationFactor float64
	Multiplier          float64
	MaxInterval         time.Duration
	MaxElapsedTime      time.Duration
}

// Resets an exponential backoff policy to its initial state.
func (backoff *exponentialBackoff) Reset() {
	backoff.currentInterval = backoff.initialInterval
}
