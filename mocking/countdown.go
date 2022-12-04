package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countStarter = 3
	finalWord    = "Go!"
)

type Sleeper interface {
	Sleep()
}

func Countdown(out io.Writer, s Sleeper) {
	for i := countStarter; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.SleepFunc(cs.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"
