package transcoder

import (
	"testing"
	"time"
)

func TestConstantBackoffReturnsConstantInterval(t *testing.T) {
	backoff := NewConstantBackoff(ConstantBackoffPolicy{
		Interval: 10 * time.Second,
	})

	got := backoff.Next()
	if got != 10*time.Second {
		t.Errorf("Next() = %v; want %v", got, 10*time.Second)
	}

	got = backoff.Next()
	if got != 10*time.Second {
		t.Errorf("Next() = %v; want %v", got, 10*time.Second)
	}
}

func TestConstantBackoffReturnsSameIntervalAfterReset(t *testing.T) {
	backoff := NewConstantBackoff(ConstantBackoffPolicy{
		Interval: 10 * time.Second,
	})

	before := backoff.Next()
	backoff.Reset()
	after := backoff.Next()

	if after != before {
		t.Errorf("Next() = %v after Reset(); %v before Reset()", after, before)
	}
}
