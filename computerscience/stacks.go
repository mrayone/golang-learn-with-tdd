package computerscience

import "fmt"

type ArrayStackOfStrings struct {
	items []string
	n     int
}

func NewArrayOfStrings(capacity int) ArrayStackOfStrings {
	return ArrayStackOfStrings{
		items: make([]string, capacity),
		n:     0,
	}
}

func (a *ArrayStackOfStrings) IsEmpty() bool {
	return a.n == 0
}

func (a *ArrayStackOfStrings) Pop() string {
	val := a.items[a.n-1]
	a.n--
	return val
}

func (a *ArrayStackOfStrings) Push(v string) {
	a.items[a.n] = v
	a.n++
}

type node struct {
	item string
	next *node
}

type LinkedStackOfStrings struct {
	first *node
}

func (a *LinkedStackOfStrings) IsEmpty() bool {
	return a.first == nil
}

func (a *LinkedStackOfStrings) Pop() string {
	item := a.first.item
	a.first = a.first.next
	return item
}

func (a *LinkedStackOfStrings) Push(v string) {
	oldFirst := a.first
	first := node{}
	first.item = v
	first.next = oldFirst
	a.first = &first
}

func (a *LinkedStackOfStrings) ToString() string {
	itens := ""
	for el := a.first; el != nil; el = el.next {
		itens += el.item + " "
	}

	return itens
}

type ResizingtackOfStrings struct {
	items []string
	n     int
}

func NewResizingtackOfStrings() ResizingtackOfStrings {
	return ResizingtackOfStrings{
		items: make([]string, 1),
		n:     0,
	}
}

func (a *ResizingtackOfStrings) resize(capacity int) {
	temp := make([]string, capacity)
	for i := 0; i < a.n; i++ {
		temp[i] = a.items[i]
	}
	a.items = temp
}

func (a *ResizingtackOfStrings) IsEmpty() bool {
	return a.n == 0
}

func (a *ResizingtackOfStrings) Pop() string {
	// Remove and return most recently inserted item.
	val := a.items[a.n-1]
	a.n--
	a.items[a.n] = "" // Avoid loitering.
	size := len(a.items)
	if a.n > 0 && a.n == size/4 {
		a.resize(size / 2)
	}

	return val
}

func (a *ResizingtackOfStrings) Push(v string) {
	// Insert item onto stack.
	size := len(a.items)
	if a.n == size {
		a.resize(2 * size)
	}

	a.items[a.n] = v
	a.n++
}

func (a *ResizingtackOfStrings) ToString() string {
	return fmt.Sprintf("%v", a.items)
}
