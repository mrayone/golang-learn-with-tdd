package computerscience_test

import (
	"fmt"

	"github.com/mrayone/learn-go/computerscience"
)

func ExampleQueueEnqueue() {
	queue := computerscience.Queue{}

	queue.Enqueue("item 1")
	queue.Enqueue("item 2")
	queue.Enqueue("item 3")

	fmt.Println(queue.ToString())
	//Output: item 1 item 2 item 3
}

func ExampleQueueDequeue() {
	queue := computerscience.Queue{}

	queue.Enqueue("item 1")
	queue.Enqueue("item 2")
	queue.Enqueue("item 3")

	queue.Dequeue() // item 1
	queue.Dequeue() // item 2
	queue.Enqueue("item 4")
	fmt.Println(queue.ToString())
	//Output: item 3 item 4
}
