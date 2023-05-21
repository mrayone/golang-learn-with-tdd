package bigonotation

import "fmt"

// Moves are baskett moves
// It is a exponencial(2^n) time in the order of growth
func Moves(n int, enter bool) {
	if n == 0 {
		return
	}

	Moves(n-1, true)
	if enter {
		fmt.Printf("enter %d\n", n)
	} else {
		fmt.Printf("exit %d\n", n)
	}
	Moves(n-1, false)
}
