package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Servers []Server
	Users   []string
	Groups  []string
}

type Server struct {
	HostName string
	Backends []string
	Loads    []float32
}

var config = &Configuration{}

func readConfigFile() {
	file, err := os.Open(*configFile)

	if err != nil {
		log.Fatal("ERROR: Could not open config file! \n", err.Error())
	}

	decoder := json.NewDecoder(file)
	decoder.Decode(&config)
	fmt.Println(config.Servers) //debug
}
