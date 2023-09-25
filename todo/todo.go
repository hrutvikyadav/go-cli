package todo

import (
	"encoding/json"
	"os"
)

// exported type from todo package
type Item struct {
	Title string
}

func SaveItems(filename string, list []Item) error {
	b, err := json.Marshal(list)
	if err != nil {
		return err
	}
	//fmt.Println(string(b))
	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var list []Item
	if err = json.Unmarshal(b, &list); err != nil {
		return []Item{}, err
	}
	return list, nil
}
