package main

import (
	"net/http"
	"os"
	"time"

	"github.com/mrayone/learn-go/di"
	"github.com/mrayone/learn-go/mocking"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
