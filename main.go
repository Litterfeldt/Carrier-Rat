package main

func main() {
	// Begin by parsing the command line arguments
	parseFlags()

	// Print the startup message
	printStartupMessage()

	// readConfigFile
	readConfigFile()

	// Finally start the load balancer
	//startBalancer()
}
