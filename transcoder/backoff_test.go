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

func TestConstantBackoffDefaultsToTenSeconds(t *testing.T) {
	backoff := NewConstantBackoff(ConstantBackoffPolicy{})
	if got := backoff.Next(); got != 10*time.Second {
		t.Errorf("Default initial Next() = %d; want %d", got, 10*time.Second)
	}
}

func TestConstantBackoffReturnsStopAfterMaxElapsedTime(t *testing.T) {
	backoff := NewConstantBackoff(ConstantBackoffPolicy{
		Interval:       10 * time.Second,
		MaxElapsedTime: 20 * time.Second,
	})

	if got := backoff.Next(); got != 10*time.Second {
		t.Errorf("First Next() = %d; want %d", got, 10*time.Second)
	}
	if got := backoff.Next(); got != 10*time.Second {
		t.Errorf("Second Next() = %d; want %d", got, 10*time.Second)
	}
	if got := backoff.Next(); got != Stop {
		t.Errorf("Next() after MaxElapsedTime = %d; want %d", got, 10*time.Second)
	}
}

func TestConstantBackoffWorksAfterReset(t *testing.T) {
	backoff := NewConstantBackoff(ConstantBackoffPolicy{
		Interval:       10 * time.Second,
		MaxElapsedTime: 10 * time.Second,
	})

	backoff.Next()
	if got := backoff.Next(); got != Stop {
		t.Errorf("Next() after MaxElapsedTime = %d; want %d", got, 10*time.Second)
	}

	backoff.Reset()

	if got := backoff.Next(); got != 10*time.Second {
		t.Errorf("Next() = %d; want %d", got, 10*time.Second)
	}
}
