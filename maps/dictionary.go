package maps

type Dictionary map[string]string

type DictionaryErr string

func (dictionaryErr DictionaryErr) Error() string {
	return string(dictionaryErr)
}

const (
	ErrNotFound     = DictionaryErr("could not find request word")
	ErrWordConflict = DictionaryErr("word to be added already exists")
)

func (dictionaryToSearch Dictionary) Search(wordToSearchFor string) (definition string, err error) {
	definition, wordFound := dictionaryToSearch[wordToSearchFor]

	if !wordFound {
		return "", ErrNotFound
	}

	return definition, nil
}

func (dictionaryToAddTo Dictionary) Add(word, definition string) (err error) {

	_, searchErr := dictionaryToAddTo.Search(word)

	switch searchErr {
	case ErrNotFound:
		dictionaryToAddTo[word] = definition
	case nil:
		return ErrWordConflict
	default:
		return searchErr
	}

	return
}

func (dictionary Dictionary) Update(word, newDefinition string) (err error) {

	_, searchErr := dictionary.Search(word)

	switch searchErr {
	case nil:
		dictionary[word] = newDefinition
	default:
		return searchErr
	}

	return
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)
}
