package concurrency

import "fmt"

func MergeSort(data []int) []int {
	if len(data) <= 1 { // stop condition
		return data
	}
	done := make(chan bool)
	defer close(done)

	mid := len(data) / 2
	var left []int
	go func() {
		left = MergeSort(data[:mid])
		done <- true
	}()
	right := MergeSort(data[mid:])
	<-done
	return merge(left, right)
}

func MergeSortTwo(data []int) []int {
	if len(data) <= 1 { // stop condition
		return data
	}

	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:] // pop first
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func Merge() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("%v\n%v\n", data, MergeSort(data))
}

func MergeSortSequence(elems ...int) {
	myChannel := make(chan int)
	defer close(myChannel)
	go func() {
		sorted := MergeSort(elems)
		for i := range sorted {
			myChannel <- sorted[i]
		}
	}()

	sorted := make([]int, 0)
	for i := 0; i < len(elems); i++ {
		sorted = append(sorted, <-myChannel)
	}

	fmt.Printf("%v\n%v\n", elems, sorted)
}

func MergeSortSequenceTwo(elems ...int) {
	fmt.Printf("%v\n%v\n", elems, MergeSortTwo(elems))
}
