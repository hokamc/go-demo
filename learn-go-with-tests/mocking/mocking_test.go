package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (sleeper *SpySleeper) Sleep() {
	sleeper.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	sleeper := &SpySleeper{}
	Countdown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!
`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if sleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", sleeper.Calls)
	}
}
