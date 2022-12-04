package maps

const (
	//Sentinel errors
	ErrorNotFoundWord   = DictionaryErr("could not find the word you were looking for")
	ErrWordExisting     = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if val, ok := d[key]; ok {
		return val, nil
	}

	return "", ErrorNotFoundWord
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFoundWord:
		d[word] = definition
	case nil:
		return ErrWordExisting
	default:
		return nil
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err != nil {
		return ErrWordDoesNotExist
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
