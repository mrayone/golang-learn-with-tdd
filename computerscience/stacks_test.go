package computerscience_test

import (
	"fmt"

	"github.com/mrayone/learn-go/computerscience"
)

func ExampleNewArrayOfStrings() {
	a := computerscience.NewArrayOfStrings(3)
	a.Push("a")
	a.Push("b")
	a.Push("c")
	fmt.Println(a.Pop())
	//Output: c
}

func ExampleLinkedStackOfStrings() {
	list := computerscience.LinkedStackOfStrings{}
	list.Push("to")
	list.Push("be")
	list.Push("not")
	list.Push("that")
	list.Push("or")
	list.Push("be")

	fmt.Println(list.ToString())
	//Output: be or that not be to
}

func ExampleNewResizingtackOfStrings() {
	res := computerscience.NewResizingtackOfStrings()
	res.Push("to")
	res.Push("be")
	res.Push("or")
	res.Push("not")
	res.Push("to")

	res.Pop()

	res.Push("be")
	res.Pop()
	res.Pop()
	res.Push("that")
	res.Pop()
	res.Pop()
	res.Pop()
	res.Push("is")

	fmt.Println(res.ToString())
	//Output:
	//[to is]
}
