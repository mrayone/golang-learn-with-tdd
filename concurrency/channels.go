package concurrency

import "fmt"

func sendValues(myIntChannel chan int) {
	for i := 0; i < 5; i++ {
		myIntChannel <- i // sending value
	}
}

func sendValuesAndClose(myIntChannel chan int) {
	for i := 0; i < 5; i++ {
		myIntChannel <- i // sending value
	}

	close(myIntChannel)
}

func Run() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 5; i++ {
		fmt.Println(<-myIntChannel)
	}
}

func RunWithDeadlock() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 6; i++ {
		fmt.Println(<-myIntChannel)
	}
}

func RunCloseAndChannel() {
	myIntChannel := make(chan int)

	go sendValuesAndClose(myIntChannel)

	for i := 0; i < 6; i++ {
		fmt.Println(<-myIntChannel)
	}
}

func RunCloseAndChannelBreak() {
	myIntChannel := make(chan int)

	go sendValuesAndClose(myIntChannel)

	for i := 0; i < 6; i++ {
		value, open := <-myIntChannel
		if !open {
			break
		}
		fmt.Println(value)
	}
}

func RunAlongRange() {
	myIntChannel := make(chan int)

	go sendValuesAndClose(myIntChannel)

	for v := range myIntChannel {
		fmt.Println(v)
	}
}
