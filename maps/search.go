package maps

import "errors"

var ErrorNotFoundWord = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func Search(dictionary Dictionary, key string) (string, error) {
	if val, ok := dictionary[key]; ok {
		return val, nil
	}

	return "", ErrorNotFoundWord
}
