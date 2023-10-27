package concurrencypatterns

import (
	"fmt"
	"math/rand"
	"time"
)

type CookInfo struct {
	foodCooked     string
	waitForPartner chan bool
}

func cookFood(name string) <-chan CookInfo {
	cookChannel := make(chan CookInfo)
	wait := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			cookChannel <- CookInfo{fmt.Sprintf("%s %s", name, "Done"), wait}
			time.Sleep(time.Duration(rand.Int31n(1e3)) * time.Millisecond)
			<-wait
		}
	}()

	return cookChannel
}

func fanInCook(ch1, ch2 <-chan CookInfo) <-chan CookInfo {
	chMerge := make(chan CookInfo)

	go func() {
		for {
			chMerge <- <-ch1
		}
	}()

	go func() {
		for {
			chMerge <- <-ch2
		}
	}()

	return chMerge
}

func CookFood() {
	chGame := fanInCook(cookFood("player 1"), cookFood("player 2"))

	for round := 0; round < 3; round++ {
		foodOne := <-chGame
		fmt.Println(foodOne.foodCooked)

		foodTwo := <-chGame
		fmt.Println(foodTwo.foodCooked)

		foodOne.waitForPartner <- true
		foodTwo.waitForPartner <- true

		fmt.Printf("Done with round %d\n", round+1)
	}

	fmt.Println("Done with the competition.")
}
