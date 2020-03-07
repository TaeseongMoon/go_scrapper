package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFount   = errors.New("Not Found")
	errWordExists = errors.New("already exist word")
	errCantUpdate = errors.New("Can't update non-existing word")
)

// Search words
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFount
}

// Add words
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFount {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

// Update words
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFount:
		return errCantUpdate
	}
	return nil
}

// Delete the word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
