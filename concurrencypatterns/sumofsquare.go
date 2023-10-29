package concurrencypatterns

import (
	"fmt"
)

func SumOfSquares(c, quit chan int) {
	side := 0
	for {
		select {
		case c <- side * side:
			side++
		case <-quit:
			return
		}
	}
}

func RunSumOfSquares() {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 0; i < 6; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 0

	}()
	SumOfSquares(mychannel, quitchannel)
}
