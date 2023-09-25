package todo

import (
	"encoding/json"
	"fmt"
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
