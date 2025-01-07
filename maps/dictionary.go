package maps

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrWordNotFound     = DictionaryErr("word not found")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, ok := d[word]
	if !ok {
		return ErrWordDoesNotExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	if _, ok := d[word]; !ok {
		return ErrWordDoesNotExist
	}
	delete(d, word)
	return nil
}
