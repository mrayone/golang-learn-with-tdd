package computerscience_test

import (
	"fmt"

	"github.com/mrayone/learn-go/computerscience"
)

func ExampleArrayOfString() {
	a := computerscience.NewArrayOfStrings(3)
	a.Push("a")
	a.Push("b")
	a.Push("c")
	fmt.Println(a.Pop())
	//Output: c
}
