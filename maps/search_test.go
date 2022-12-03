package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		desc          string
		dictionary    Dictionary
		word          string
		expectedText  string
		expectedError error
	}{
		{
			desc:         "known word",
			word:         "test",
			dictionary:   Dictionary{"test": "this is just a test"},
			expectedText: "this is just a test",
		},
		{
			desc:          "unknown word",
			word:          "unknown",
			dictionary:    Dictionary{"test": "this is just a test"},
			expectedError: ErrorNotFoundWord,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := Search(tC.dictionary, tC.word)

			assertError(t, err, tC.expectedError)
			assertStrings(t, got, tC.expectedText)
		})
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q given", got, want)
	}
}
