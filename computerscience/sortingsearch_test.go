package computerscience_test

import (
	"bytes"
	"fmt"
	"reflect"
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

type TestCaseSearch[T comparable] struct {
	desc     string
	key      T
	values   []computerscience.Comparable[T]
	expected int
}

func TestSearch(t *testing.T) {
	stackStrings := []computerscience.Comparable[computerscience.String]{
		computerscience.String("aaa"),
		computerscience.String("bbb"),
		computerscience.String("ccc"),
		computerscience.String("ooo"),
	}

	testCasesStrings := []TestCaseSearch[computerscience.String]{
		{
			desc:     "find key aaa",
			values:   stackStrings,
			key:      "aaa",
			expected: 0,
		},
		{
			desc:     "find key bbb",
			values:   stackStrings,
			key:      "bbb",
			expected: 1,
		},
		{
			desc:     "find key ooo",
			values:   stackStrings,
			key:      "ooo",
			expected: 3,
		},
		{
			desc:     "find key ccc",
			values:   stackStrings,
			key:      "ccc",
			expected: 2,
		},
	}
	for _, tc := range testCasesStrings {
		t.Run(tc.desc, func(t *testing.T) {
			got := computerscience.Search(tc.key, tc.values)

			if got != tc.expected {
				t.Errorf("unexpected value, got %d want %d", got, tc.expected)
			}
		})
	}

	personsStack := []computerscience.Comparable[computerscience.Person]{
		computerscience.Person{Name: "Alisson Doe"},
		computerscience.Person{Name: "Bob Doe"},
		computerscience.Person{Name: "John Doe"},
		computerscience.Person{Name: "Joana Doe"},
	}

	testCasesPerson := []TestCaseSearch[computerscience.Person]{
		{
			desc:     "not find person Maycon",
			values:   personsStack,
			key:      computerscience.Person{Name: "Maycon"},
			expected: -1,
		},
		{
			desc:     "find key Bob Doe",
			values:   personsStack,
			key:      computerscience.Person{Name: "Bob Doe"},
			expected: 1,
		},
		{
			desc:     "find key John Doe",
			values:   personsStack,
			key:      computerscience.Person{Name: "John Doe"},
			expected: 2,
		},
	}
	for _, tc := range testCasesPerson {
		t.Run(tc.desc, func(t *testing.T) {
			got := computerscience.Search(tc.key, tc.values)

			if got != tc.expected {
				t.Errorf("unexpected value, got %d want %d", got, tc.expected)
			}
		})
	}
}

type TestCaseSort[T comparable] struct {
	desc     string
	values   []computerscience.Comparable[T]
	expected []computerscience.Comparable[T]
}

func TestMergeSort(t *testing.T) {
	stack := []computerscience.Comparable[computerscience.Int]{
		computerscience.Int(7),
		computerscience.Int(2),
		computerscience.Int(3),
		computerscience.Int(4),
		computerscience.Int(1),
		computerscience.Int(5),
		computerscience.Int(6),
	}

	testCasesStrings := []TestCaseSort[computerscience.Int]{
		{
			desc:   "sort array",
			values: stack,
			expected: []computerscience.Comparable[computerscience.Int]{
				computerscience.Int(1),
				computerscience.Int(2),
				computerscience.Int(3),
				computerscience.Int(4),
				computerscience.Int(5),
				computerscience.Int(6),
				computerscience.Int(7),
			},
		},
	}
	for _, tc := range testCasesStrings {
		t.Run(tc.desc, func(t *testing.T) {
			computerscience.MergeSort(tc.values)

			if !reflect.DeepEqual(tc.values, tc.expected) {
				t.Errorf("unexpected value, got %d want %d", tc.values, tc.expected)
			}
		})
	}
}

func ExampleInsertionSort() {
	personsStack := []computerscience.Comparable[computerscience.Person]{
		computerscience.Person{Name: "John Doe"},
		computerscience.Person{Name: "Alisson Doe"},
		computerscience.Person{Name: "Joana Doe"},
		computerscience.Person{Name: "Bob Doe"},
	}
	computerscience.InsertionSort(personsStack)
	fmt.Println(personsStack)
	//Output: [{Alisson Doe  0} {Bob Doe  0} {Joana Doe  0} {John Doe  0}]

}
