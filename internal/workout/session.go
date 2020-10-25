package workout

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"time"
)

// Session holds data for a workout.
type Session struct {
	Sets []Set
}

// Dictate announces a session.
func (s *Session) Dictate() {
	speak("starting session in 10")
	time.Sleep(10 * time.Second)

	for i, set := range s.Sets {
		set.dictate(i)
	}

	speak("done!")
}

const timeDelay = time.Second

// Set holds data for a set within a session.
type Set struct {
	// Dur is the overall length of the set.
	Dur time.Duration

	// Coundown is the length of time to announce
	// each second before the set starts.
	Countdown time.Duration

	// Countup is the length of time to accounce
	// each second after the set starts.
	Countup time.Duration
}

// Dictate announces a set.
func (s *Set) dictate(i int) {
	// Accounce the start of the set.
	msg := fmt.Sprintf("set %v", i+1)
	speak(msg)
	time.Sleep(time.Second)

	// Accounce the 'countup'.
	secs := int(math.Round(s.Countup.Seconds()))
	for i := 1; i <= secs; i++ {
		msg := strconv.Itoa(i)
		speak(msg)
		time.Sleep(timeDelay)
	}

	// Wait for the set duration less the
	// countdown.
	d := s.Dur - s.Countdown - s.Countup
	timer := time.NewTimer(d)
	<-timer.C

	// Accounce the countdown.
	secs = int(math.Round(s.Countdown.Seconds()))
	for i := secs; i > 0; i-- {
		msg := strconv.Itoa(i)
		speak(msg)

		time.Sleep(timeDelay)
	}
}

const cmdSpeak = "espeak"

func speak(msg string) {
	cmd := exec.Command(cmdSpeak, msg)
	go cmd.Run()
}
