package concurrencypatterns

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Race(channel, quit chan string, i int) {

	channel <- fmt.Sprintf("Car %d started!", i)
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		quit <- fmt.Sprintf("Car %d reached the finishing line!", i)
		fmt.Println(<-quit) // trava a execução até receber do sender
		wg.Done()           // seta como done o waiting group
	}
}

func RunRace() {

	channel := make(chan string)
	quit := make(chan string)
	wg.Add(1) // irá esperar por 1 goroutine
	for i := 0; i < 3; i++ {
		go Race(channel, quit, i)
	}

	for {
		select {
		case raceUpdates := <-channel:
			fmt.Println(raceUpdates)
		case winnerAnnoucement := <-quit:
			fmt.Println(winnerAnnoucement)
			quit <- "You win!" // envia o you win, e trava a execução até o receiver
			wg.Wait()          // após receber o sinal de done, libera o for.
			return

		}
	}
}
