package concurrency_test

import "github.com/mrayone/learn-go/concurrency"

func ExampleMerge() {
	concurrency.Merge()
	//Output:
	//[9 4 3 6 1 2 10 5 7 8]
	//[1 2 3 4 5 6 7 8 9 10]
}

func ExampleMergeSortSequence() {
	concurrency.MergeSortSequence(9, 4, 3, 6, 1, 2, 10, 5, 7, 8)
	//Output:
	//[9 4 3 6 1 2 10 5 7 8]
	//[1 2 3 4 5 6 7 8 9 10]
}

func ExampleMergeSortSequenceTwo() {
	concurrency.MergeSortSequenceTwo(9, 4, 3, 6, 1, 2, 10, 5, 7, 8)
	//Output:
	//[9 4 3 6 1 2 10 5 7 8]
	//[1 2 3 4 5 6 7 8 9 10]
}
