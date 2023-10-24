package config

import (
	"encoding/json"
	"os"
)

type config struct {
	Base    int      `json:"alphabet_base"`
	Start   byte     `json:"alphabet_start"`
	Len     int      `json:"password_len"`
	Threads int      `json:"threads"`
	Cases   []string `json:"cases"`
}

var Config config

func init() {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		panic(err)
	}
}
