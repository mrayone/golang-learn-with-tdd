package computerscience

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
