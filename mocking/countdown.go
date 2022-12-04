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

type DefaultSleeper struct{}

func (ds *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
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
