package computerscience_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mrayone/learn-go/computerscience"
)

func TestPlayTweentyQuestions(t *testing.T) {
	var mockReader bytes.Buffer

	testCases := []struct {
		desc          string
		inputs        []string
		expectedValue string
		expectedErr   error
	}{
		{
			desc: "guess the 77 number",
			inputs: []string{
				fmt.Sprintf("%d\n", 7), // questions quantity
				"true\n",
				"false\n",
				"false\n",
				"true\n",
				"true\n",
				"false\n",
				"true\n",
			},
			expectedValue: fmt.Sprintf("Your number is %d\n", 77),
		},
		{
			desc: "guess the 44 number",
			inputs: []string{
				fmt.Sprintf("%d\n", 6), // questions quantity 101100
				"true\n",
				"false\n",
				"true\n",
				"true\n",
				"false\n",
				"false\n",
			},
			expectedValue: fmt.Sprintf("Your number is %d\n", 44),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			for _, val := range tc.inputs {
				mockReader.Write([]byte(val))
			}

			guess, err := computerscience.PlayTwentyQuestions(&mockReader)

			assertError(t, err, tc.expectedErr)

			if guess != tc.expectedValue {
				t.Errorf("unexpected value %q given %q want", guess, tc.expectedValue)
			}
		})
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q given", got, want)
	}
}
