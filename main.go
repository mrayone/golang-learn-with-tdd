package main

import (
	"net/http"
	"os"

	"github.com/mrayone/learn-go/di"
	"github.com/mrayone/learn-go/mocking"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))

	mocking.Countdown(os.Stdout, &mocking.DefaultSleeper{})
}
