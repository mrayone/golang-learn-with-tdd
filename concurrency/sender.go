package concurrency

import (
	"fmt"
	"time"
)

func Sender() {
	tsk := make(chan bool)

	go func(task chan bool) {
		fmt.Println("sleeping 5s")
		time.Sleep(5 * time.Second)

		task <- true
		close(task)
	}(tsk)

	fmt.Println("logging")

	go func(task chan bool) {
		for {
			select {
			case <-tsk:
				fmt.Println("log work")
				return
			}
		}
	}(tsk)
}
