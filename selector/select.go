package selector

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimedout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimedout)
}

// struct{} is the smallest data type available from a memory perspective
// since we are closing and not sending anything on the chan.
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// <- blocking call, as I'm waiting for a value
	// select lets you do is wait on multiple channels. The firsto one to send a value "wins"
	// and the code underneath case is executed.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // it's perfect to scape of code that can block forever de channels
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// func meansureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
