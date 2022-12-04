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
			got, err := tC.dictionary.Search(tC.word)

			assertError(t, err, tC.expectedError)
			assertStrings(t, got, tC.expectedText)
		})
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		desc        string
		dictionary  Dictionary
		word        string
		definition  string
		expectedErr error
	}{
		{
			desc:       "add a word",
			dictionary: Dictionary{},
			word:       "test",
			definition: "this is just a test",
		},
		{
			desc:        "existing word",
			dictionary:  Dictionary{"test": "existing"},
			word:        "test",
			definition:  "existing",
			expectedErr: ErrWordExisting,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.dictionary.Add(tC.word, tC.definition)
			assertError(t, err, tC.expectedErr)
			assertDefinition(t, tC.dictionary, tC.word, tC.definition)
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		dictionary  Dictionary
		word        string
		definition  string
		expectedErr error
	}{
		{
			desc:       "update a word",
			dictionary: Dictionary{"test": "existing"},
			word:       "test",
			definition: "this is just a test",
		},
		{
			desc:        "update unexisting word",
			dictionary:  Dictionary{"test": "existing"},
			expectedErr: ErrWordDoesNotExist,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.dictionary.Update(tC.word, tC.definition)
			assertError(t, err, tC.expectedErr)
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		desc        string
		dictionary  Dictionary
		word        string
		expectedErr error
	}{
		{
			desc:       "delete existing value",
			dictionary: Dictionary{"test": "test"},
			word:       "test",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.dictionary.Delete(tC.word)
		})
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
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
