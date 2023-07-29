package computerscience

import "fmt"

type nodeQueue struct {
	item any
	next *nodeQueue
}

type Queue struct {
	first *nodeQueue
	last  *nodeQueue
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Enqueue(item any) {
	oldLast := q.last
	q.last = &nodeQueue{}
	q.last.item = item
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
}

func (q *Queue) Dequeue() any {
	item := q.first.item
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	return item
}

func (a *Queue) ToString() string {
	itens := ""
	for el := a.first; el != nil; el = el.next {
		itens += fmt.Sprintf("%v ", el.item)
	}

	return itens
}
