package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Items []DataItem
}

type DataItem struct {
	Timezone string   `json:"timezone"`
	Payload  []string `json:"payload"`
	Tag      string   `json:"tag"`
}

func (data *Data) Get() {
	file, err := os.Open(FilePath("data.json"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data.Items)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
