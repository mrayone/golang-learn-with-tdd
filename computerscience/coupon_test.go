package computerscience_test

import (
	"fmt"
	"testing"

	"github.com/mrayone/learn-go/computerscience"
)

func TestCollector(t *testing.T) {
	testCases := []struct {
		desc  string
		input int
	}{
		{
			desc:  "input 1000 value",
			input: 1000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			output := computerscience.Collector(tc.input)

			fmt.Println(output)
			if output == 0 {
				t.Errorf("unexpected output value %d", output)
			}
		})
	}
}
