package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	testCases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two fields",
			Input: struct {
				Name, City string
			}{
				"Chris",
				"London",
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "nested fields",
			Input: struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
		{
			Name: "map",
			Input: map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			ExpectedCalls: []string{"Bar", "Boz"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var got []string
			walk(tc.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tc.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tc.ExpectedCalls)
			}
		})
	}
}
