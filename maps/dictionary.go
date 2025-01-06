package maps

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrKeyNotFound = DictionaryErr("key not found")
	ErrKeyExists   = DictionaryErr("key already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if ok {
		return ErrKeyExists
	}
	d[key] = value
	return nil
}
