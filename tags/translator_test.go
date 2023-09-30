package tags_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/mrayone/learn-go/tags"
	"testing"
)

func TestGetOrigins(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    tags.Config
		expectedOrigns map[string]string
	}{
		{
			name:        "should return correct origins",
			inputStruct: tags.Config{},
			expectedOrigns: map[string]string{
				"settings.canPlay":         "allowPlay",
				"settings.canWatchVideo":   "allowWatchVideo",
				"settings.priceMovie":      "priceOfMovie",
				"settings.discountEnabled": "toggleDiscount",
				"settings.defaultMessage":  "message",
				"settings.value":           "objectValue",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.inputStruct.GetOrigins()
			if diff := cmp.Diff(got, tc.expectedOrigns); diff != "" {
				t.Errorf("unexpected output diff: %s", diff)
			}
		})
	}
}

//BenchmarkFindAllConfig-16         835246              1396 ns/op             720 B/op         10 allocs/op
//BenchmarkFindAllMaps-16          4517833               266.5 ns/op             0 B/op          0 allocs/op

func TestFindAllConfig(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected map[string]string
	}{
		{
			name: "should find data correctly",
			input: map[string]string{
				"settings.canPlay":       "true",
				"settings.canWatchVideo": "true",
			},
			expected: map[string]string{
				"allowPlay":       "true",
				"allowWatchVideo": "true",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tags.FindAllConfig(tc.input)
			if err != nil {
				t.Errorf("unexpected error %s", err)
			}

			if diff := cmp.Diff(got, tc.expected); diff != "" {
				t.Errorf("unexpected output diff: %s", diff)
			}
		})
	}
}

func BenchmarkFindAllConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tags.FindAllConfig(map[string]string{
			"settings.canPlay":        "true",
			"settings.canWatchVideo":  "true",
			"settings.canDoSomething": "true",
			"settings.priceMovie":     "22.6",
			"settings.defaultMessage": "Where are u java?",
		})

	}
}

func BenchmarkFindAllMaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tags.FindAllWithMaps(map[string]string{
			"settings.canPlay":        "true",
			"settings.canWatchVideo":  "true",
			"settings.canDoSomething": "true",
			"settings.priceMovie":     "22.6",
			"settings.defaultMessage": "Where are u java?",
		}, map[string]string{
			"settings.canPlay":        "allowPlay",
			"settings.canWatchVideo":  "allowWatchVideo",
			"settings.canDoSomething": "toggleOn",
			"settings.priceMovie":     "priceOfMovie",
			"settings.defaultMessage": "message",
		})
	}
}
