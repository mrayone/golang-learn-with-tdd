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
