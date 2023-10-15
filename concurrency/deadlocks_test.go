package concurrency_test

import (
	"github.com/mrayone/learn-go/concurrency"
	"testing"
)

func TestDeadlock(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "shoulds display deadlock",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			concurrency.Deadlock()
		})
	}
}

func ExampleAvoidChannelBlock() {
	concurrency.AvoidChannelBlock()
	//Output: 10
}

func ExampleAvoidChannelBlockTwo() {
	concurrency.AvoidChannelBlockTwo()
	//Output: 20
}
