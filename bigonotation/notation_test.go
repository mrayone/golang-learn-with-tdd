package bigonotation_test

import (
	"testing"

	"github.com/mrayone/learn-go/bigonotation"
)

func TestConstantTime(t *testing.T) {
	testCases := []struct {
		desc     string
		slice    []string
		expected string
	}{
		{
			desc: "correctly value",
			slice: []string{
				"abced",
			},
			expected: "abced",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := bigonotation.ConstantTime(tc.slice)

			if err != nil {
				t.Error("unexpected error", err)
			}

			if got != tc.expected {
				t.Errorf("unexpected value, want %s and got %s", tc.expected, got)
			}
		})
	}
}

func TestLinearTime(t *testing.T) {
	testCases := []struct {
		desc     string
		slice    []string
		expected string
	}{
		{
			desc: "correctly value",
			slice: []string{
				"g",
				"o",
				"l",
				"a",
				"ng",
			},
			expected: "golang",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := bigonotation.LinearTime(tc.slice)

			if err != nil {
				t.Error("unexpected error", err)
			}

			if got != tc.expected {
				t.Errorf("unexpected value, want %s and got %s", tc.expected, got)
			}
		})
	}
}
