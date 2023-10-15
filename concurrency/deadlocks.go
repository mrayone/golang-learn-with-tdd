package concurrency

import "fmt"

func Deadlock() {
	mychannel := make(chan int)
	mychannel <- 10
	fmt.Println(<-mychannel)
}

func AvoidChannelBlock() {
	mychannel := make(chan int)

	go func() {
		mychannel <- 10
	}()

	fmt.Println(<-mychannel)
}

func AvoidChannelBlockTwo() {
	mychannel1 := make(chan int)
	mychannel2 := make(chan int)
	mychannel3 := make(chan int)
	go func() {
		<-mychannel1
	}()

	go func() {
		mychannel2 <- 20
	}()

	go func() {
		<-mychannel3
	}()

	fmt.Println(<-mychannel2)
}
