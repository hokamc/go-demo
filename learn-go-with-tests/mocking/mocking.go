package mocking

import (
	"bytes"
	"fmt"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

func (sleeper DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out *bytes.Buffer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	fmt.Fprintln(out, finalWord)
}
