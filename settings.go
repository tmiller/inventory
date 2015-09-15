package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type settings struct {
	ServerPort      string `json:"ServerPort"`
	ForemanURL      string `json:"ForemanURL"`
	ForemanUsername string `json:"ForemanUsername"`
	ForemanPassword string `json:"ForemanPassword"`
}

func loadSettings(filename string) settings {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error:", err)
	}

	var settings settings
	err = json.Unmarshal(content, &settings)
	if err != nil {
		log.Println("Error:", err)
	}

	return settings
}

func (s settings) String() string {
	return fmt.Sprintf("ServerPort: %s\nForemanUsername: %s\n", s.ServerPort, s.ForemanUsername)
}
