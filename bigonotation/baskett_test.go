package bigonotation_test

import "github.com/mrayone/learn-go/bigonotation"

func ExampleMoves() {
	arg := 2
	bigonotation.Moves(arg, true)
	//Output:
	//enter 1
	//enter 2
	//exit 1
}
