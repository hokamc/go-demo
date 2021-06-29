package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrorNotFound   = DictionaryErr("could not find the word")
	ErrorWordExists = DictionaryErr("word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}
