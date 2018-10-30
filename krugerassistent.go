package main

import (
	"os"
	"fmt"
	"encoding/json"
)


type Config struct {
	TokenTelegram string `json:"TokenTelegram"`
	TittleChat string `json:"TitleChat"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func main() {
	cfg := LoadConfiguration("config.json")
	start(cfg)
}
