package concurrencypatterns

import (
	"fmt"
	"time"
)

func getNews(ch chan string) {
	newArr := []string{"Roger Federer wins the Wimbledon", "Space Exploration has reached new heights", "Wandering cat prevents playground accident"}
	for _, news := range newArr {
		ch <- news
	}
	ch <- "done"
	close(ch)
}

func ForLopRun() {
	myCh := make(chan string)

	go getNews(myCh)

	for {
		select {
		case news := <-myCh:
			fmt.Println(news)
			if news == "done" {
				return
			}
		default:
		}
	}
}

func ForLoopRun2() {
	done := make(chan string)

	for _, fruit := range []string{"apple", "banana", "cherry"} {
		select {
		case <-done:
			return
		default:
			fmt.Println(fruit)

		}
	}
}

func ForLoop3() {
	done := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 1000)
		close(done)
	}()

	for {
		select {
		case <-done: // ao fechar o canal ou receber operação, executa o return;
			fmt.Println("Work is done")
			return
		case <-time.After(3 * time.Second):
			fmt.Println("Tempo limite atigingido o trabalho ainda Não terminou.")
			return
		}
	}
}
