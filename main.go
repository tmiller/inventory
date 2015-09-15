package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var serverConfig config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var configFile string
	flag.StringVar(&configFile, "config", "./config.json", "Path to config file")
	flag.Parse()

	serverConfig, err := loadConfig(configFile)

	if os.IsNotExist(err) {
		log.Fatalf("Can't find config file: %s", configFile)
	} else if err != nil {
		log.Fatal(err)
	}

	setupRoutes()

	address := fmt.Sprintf(":%s", serverConfig.ServerPort)
	http.ListenAndServe(address, nil)
}
