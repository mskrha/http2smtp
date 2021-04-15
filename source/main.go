package main

import (
	"fmt"
	"os"
)

var (
	version string

	config configGlobal

	run bool
)

func main() {
	fmt.Printf("HTTP to SMTP proxy, version %s\n\n", version)

	/*
		Parse the configuration file
	*/
	if err := parseConfig(); err != nil {
		fmt.Println("Failed to parse the configuration file!")
		fmt.Println(err)
		return
	}

	/*
		Initialise the global variables
	*/
	run = true
	debug("Starting, PID: %d", os.Getpid())

	/*
		Prepare and start the HTTP server
	*/
	go startHttp()

	/*
		Block and wait for the system signals
	*/
	waitForSignal()
}
