package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var timeFormat = "2006-01-02 15:04:05"
var dataColumnName = "Data"
var config *Config
var once sync.Once

type Config struct {
	ShowHeaders    bool   `json:"show_headers"`
	DataColumnName string `json:"data_column_name"`
	TimeFormat     string
}

func GetConfig() *Config {
	once.Do(func() {
		// Init with default values
		config = &Config{
			ShowHeaders:    true,
			DataColumnName: dataColumnName,
			TimeFormat:     timeFormat,
		}

		parseConfigFile(config)
	})
	return config
}

func parseConfigFile(config *Config) {
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
}
