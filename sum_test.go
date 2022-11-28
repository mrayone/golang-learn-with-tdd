package main // package main's common practice to only containt ingration of other packages and not unit-testable
// for it, I rename the module in go.mod

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("create an array with length defined", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{0, 9}, []int{1, 2})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper() // this function tell to go report the line number from function call instead our assertionHelper
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of slices", func(t *testing.T) {
		got := SumAllTails([]int{0, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safaly sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

/*
When we work with big slices maybe work with copy have a good way to
collect it to the heap.
Ex: https://go.dev/play/p/Poth8JS28sc
*/
