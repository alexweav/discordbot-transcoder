package transcoder

import "time"

// Provides a back-off policy for retrying some operation.
type Backoff interface {
	Next() time.Duration
	Reset()
}

// Specifies the rate of a constant back-off policy.
type ConstantBackoffPolicy struct {
	Interval time.Duration
}

// A back-off policy where retries occur at the same constant interval.
type constantBackoff struct {
	interval time.Duration
}

// Builds a new constant back-off policy, where retries occur at the same interval.
func NewConstantBackoff(policy ConstantBackoffPolicy) Backoff {
	return &constantBackoff{
		interval: policy.Interval,
	}
}

// Gets the next back-off interval from a constant backoff.
func (backoff *constantBackoff) Next() time.Duration {
	return backoff.interval
}

// Resets a constant backoff.
func (backoff *constantBackoff) Reset() {}
