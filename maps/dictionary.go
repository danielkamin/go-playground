package main

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word")
	ErrWordExists       = DictionaryErr("word aready exists")
	ErrWordDoesNotExist = DictionaryErr("could not perform update as it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	defnition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defnition, nil
}

func (d Dictionary) Add(key, word string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = word
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, word string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = word
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}
	return nil
}
