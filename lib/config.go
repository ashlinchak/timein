package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

var timeFormat = "2006-01-02 15:04:05"

type Config struct {
	ShowHeaders bool `json:"show_headers"`
	TimeFormat  string
}

func (config *Config) SetDefaults() {
	file, err := os.Open(FilePath("config.json"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config.TimeFormat = timeFormat
}
