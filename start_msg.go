package main

import "fmt"

func printStartupMessage() {
	fmt.Println()
	fmt.Println("Carrier rat is starting up!\n")

	fmt.Printf("port:        %d\n", *port)
	fmt.Printf("cfg-file:    %s\n", *configFile)

	fmt.Println()
}
