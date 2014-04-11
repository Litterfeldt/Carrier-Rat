package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultPort       = 8080
	defaultConfigFile = "./config.json"
)

var port = flag.Int("port", defaultPort, "The port number")
var configFile = flag.String("config", defaultConfigFile, "The config file")

func parseFlags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nCommand line arguments:\n\n")
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(2)
	}

	flag.Parse()
}
