package bigonotation

// source: https://flexiple.com/algorithms/big-o-notation-cheat-sheet/

import (
	"errors"
	"fmt"
)

// ConstantTime O(1)
// When there is no dependence on the input size n, an algorithm is said to have
// a constant time of order O(1). The function bellow will require one execution step.
// whether the above array contains 1, 100 or 10000 elements. As result the function in constant time with time.
// complexity O(1)
func ConstantTime(list []string) (string, error) {
	if len(list) == 0 {
		return "", errors.New("the list should have a value")
	}

	return list[0], nil
}

// LinearTime O(n)
// Is archived when the running time of an algorithm increase linearly with
// the length of the input. This means that when a function runs for or iterates
// over an input size of n, it is said to have time complexity of order O(n)
func LinearTime(list []string) (string, error) {
	output := ""
	for _, value := range list {
		output += value
	}

	return output, nil
}

// Logarithmic Time O(log n)
// The Binary Search method takes a sorted list of elements and
// searches through it for the element x.
// With every iteration, the size of our search list shrinks by half. Therefore traversing and finding an entry in the list takes O(log(n)) time.
func BinarySearch(list []int, x int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + (high-low)/2
		if list[mid] == x {
			return mid
		}

		if list[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// QuadraticTime O(n^2)
/*
We have two nested loops in the example above. If the array has n items, the outer loop will
execute n times, and the inner loop will execute n times for each iteration of the outer loop,
resulting in n^2 prints. If the size of the array is 10, then the loop runs 10x10 times.
So the function ten will print 100 times. As a result, this function will take O(n^2) time to complete.
*/
func QuadraticTime(lst []int, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("Iteration %d, Eelement of list at %d is %d", i, j, lst[j])
		}
	}
}

// ExponentialTime: O(2^n)
// fibonacci code related
/*
With each addition to the input (n), the growth rate doubles, and the algorithm
iterates across all subsets of the input elements.
When an input unit is increased by one, the number of operations executed is doubled.

*/
func ExponentialTime(n int) int {
	if n <= 1 {
		return 1
	}

	return ExponentialTime(n-2) + ExponentialTime(n-1)
}
