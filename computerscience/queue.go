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

/*
Random queue. A random queue is a collection that supports the following
API: RandomQueue<Item>
	RandomQueue()	create an empty random queue
	boolean isEmpty() is the random queue empty?
  enqueue(Item item) add item to the random queue
  dequeue() Item remove and return a random item
						(sample without replacement)
  sample() Item return a random item, but do not remove
					(sample with replacement)

API for a generic random queue
Write a class RandomQueue that implements this API. Hint : Use a resizing array. To
remove an item, swap one at a random position (indexed 0 through n-1) with the
one at the last position (index n-1). Then, remove and return the last item, as in
ResizingArrayStack. Write a client that prints a deck of cards in random order using RandomQueue<Card>.
*/

// Random queue
