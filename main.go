package main

import (
	"fmt"
	"math"
)

// import (
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/mrayone/learn-go/di"
// 	"github.com/mrayone/learn-go/mocking"
// )

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	di.Greet(w, "world")
// }

// func main() {
// 	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
// 	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
// 	mocking.Countdown(os.Stdout, sleeper)
// }

func main() {
	fmt.Println(secondsinradians())
}

func zero() float64 {
	return 0.0
}

func secondsinradians() float64 {
	return (math.Pi / (30 / (float64(zero()))))
}
