package transcoder

import "time"

type Backoff interface {
	Next() time.Duration
	Reset()
}

type ConstantBackoffPolicy struct {
	Interval time.Duration
}

type constantBackoff struct {
	interval time.Duration
}

func NewConstantBackoff(policy ConstantBackoffPolicy) Backoff {
	return &constantBackoff{
		interval: policy.Interval,
	}
}

func (backoff *constantBackoff) Next() time.Duration {
	return backoff.interval
}

func (backoff *constantBackoff) Reset() {}
