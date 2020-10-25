package main

import (
	"log"
	"os/exec"
	"time"

	"github.com/kheppenstall/timer/internal/workout"
)

func main() {
	set := workout.Set{
		Dur: 2 * time.Minute,
		Countdown: 5 * time.Second,
		Countup: 15 * time.Second
	}

	var session workout.Session
	for i := 0; i < 20; i++ {
		sess.Sets = append(sess.Sets, set)
	}

	sess.Dictate()
}
