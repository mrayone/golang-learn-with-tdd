package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func Buzzgame() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel1 <- "Player 1 Buzzed"
	}()

	go func() {
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel2 <- "Player 2 Buzzed"
	}()

	for i := 0; i < 2; i++ {
		select {
		case p1 := <-channel1:
			fmt.Println(p1)
		case p2 := <-channel2:
			fmt.Println(p2)
		}
	}

	//fmt.Println(<-channel1)
	//fmt.Println(<-channel2)
}
