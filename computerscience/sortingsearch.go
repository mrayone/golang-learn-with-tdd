package computerscience

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
)

// search package
type Comparable[T comparable] interface {
	CompareTo(value T) int
}

func Search[T comparable](key T, arr []Comparable[T]) int {
	return search(key, arr, 0, len(arr)-1)
}

// Search consider only strings for now
func search[T comparable](key T, arr []Comparable[T], lo, hi int) int {
	if hi < lo {
		return -1
	}

	mid := (hi + lo) / 2
	comp := arr[mid].CompareTo(key)
	if comp > 0 {
		hi = mid - 1
	} else if comp < 0 {
		lo = mid + 1
	} else {
		return mid
	}

	return search(key, arr, lo, hi)
}

// PlayTwentyQuestions for given integer input
func PlayTwentyQuestions(rd io.Reader) (string, error) {
	scanner := bufio.NewScanner(rd)
	err := scanner.Err()
	if err != nil {
		return "", errors.New("you must give a number")
	}
	scanner.Scan()

	input := scanner.Text()
	k, err := strconv.Atoi(input)
	if err != nil {
		return "", errors.New("the value of input must be a valid integer")
	}

	n := int(math.Pow(2, float64(k)))
	fmt.Printf("Think of a number between 0 and %d\n", (n - 1))

	guess, err := binarySearchRecurse(0, n, scanner)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Your number is %d\n", guess), nil
}

func binarySearchRecurse(lo, hi int, scanner *bufio.Scanner) (int, error) {
	if hi-lo == 1 {
		return lo, nil
	}

	mid := lo + (hi-lo)/2
	fmt.Printf("Greater than or equal to %d? ", mid)
	err := scanner.Err()
	if err == nil {
		scanner.Scan()
		text := scanner.Text()
		isTrue, err := strconv.ParseBool(text)
		if err != nil {
			return 0, errors.New("you must given a valid boolean assertion, true or false")
		}

		if !isTrue {
			return binarySearchRecurse(lo, mid, scanner)
		} else {
			return binarySearchRecurse(mid, hi, scanner)
		}
	}

	return 0, err
}

// Real world
type String string

func (s String) CompareTo(v String) int {
	if s < v {
		return -1
	}
	if s > v {
		return 1
	}

	return 0
}

type Person struct {
	Name, Job string
	Age       int
}

func (s Person) CompareTo(v Person) int {
	if s.Name < v.Name {
		return -1
	}
	if s.Name > v.Name {
		return 1
	}

	return 0
}
