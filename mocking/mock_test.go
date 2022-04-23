package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run(
		"prints 3 to Go!", func(t *testing.T) {
			buffer := &bytes.Buffer{}
			Countdown(buffer, &SpyCountDownOperations{})

			got := buffer.String()
			want := `3
2
1
Go!`
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		},
	)

	t.Run(
		"sleep before every print", func(t *testing.T) {
			spySleepPrinter := &SpyCountDownOperations{}
			Countdown(spySleepPrinter, spySleepPrinter)

			want := []string{
				sleep,
				write,
				sleep,
				write,
				sleep,
				write,
				sleep,
				write,
			}

			if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
				t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
			}
		},
	)
}

type SpyCountDownOperations struct {
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
