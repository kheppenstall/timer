package workout

import (
	"math"
	"testing"
	"time"
)

func TestSession_Dictate(t *testing.T) {
	start := time.Now()

	sess := Session{
		Sets: []Set{
			{
				Dur:       20 * time.Second,
				Countdown: 5 * time.Second,
				Countup:   5 * time.Second,
			},
			{
				Dur:       20 * time.Second,
				Countdown: 5 * time.Second,
				Countup:   5 * time.Second,
			},
		},
	}

	sess.Dictate()

	end := time.Now()
	dur := end.Sub(start)

	// Add 5 seconds for start of session and
	// 1 for each set.
	if math.Round(dur.Seconds()) != 47 {
		t.Errorf("got %v, want %v", dur.Seconds(), 47)
	}
}
