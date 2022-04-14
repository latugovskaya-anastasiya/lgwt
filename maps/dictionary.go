package dictionary

type Dict map[string]string

func (d Dict) Search(word string) (string, error) {
	return d[word], nil
}
