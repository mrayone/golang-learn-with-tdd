package computerscience

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
)

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
