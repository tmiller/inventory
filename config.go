package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type config struct {
	ServerPort      string `json:"ServerPort"`
	ForemanURL      string `json:"ForemanURL"`
	ForemanUsername string `json:"ForemanUsername"`
	ForemanPassword string `json:"ForemanPassword"`
}

func loadConfig(filename string) (config config, err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	return
}

func (s config) String() string {
	return fmt.Sprintf("ServerPort: %s\nForemanUsername: %s\n", s.ServerPort, s.ForemanUsername)
}
