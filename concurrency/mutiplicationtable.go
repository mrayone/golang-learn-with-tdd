package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func printTable(n int, wg *sync.WaitGroup) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
	wg.Done()
}

func PrintTable() {
	var wg sync.WaitGroup

	wg.Add(12)
	for number := 1; number <= 12; number++ {
		go printTable(number, &wg)
		time.Sleep(time.Millisecond * 100)
	}

	wg.Wait()
}

func printMainTable(n int, wg *sync.WaitGroup) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, n*i)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}

func PrintMain() {
	var wg sync.WaitGroup

	for number := 2; number <= 12; number++ {
		wg.Add(1)
		go printMainTable(number, &wg)
	}

	wg.Wait()
}
