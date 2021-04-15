package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mskrha/gosmtp"
)

/*
	Show debugging output, if enabled
*/
func debug(format string, a ...interface{}) {
	if config.Debug {
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05.000000000")+": "+format+"\n", a...)
	}
}

/*
	Blocking wait for system signal
*/
func waitForSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGHUP)
	handleSignal(<-c)
	for run {
		time.Sleep(time.Second)
		handleSignal(<-c)
	}
}

/*
	Handle the system signals
*/
func handleSignal(s os.Signal) {
	switch s {
	case syscall.SIGINT:
		debug("handleSignal: Received SIGINT")
		run = false
	case syscall.SIGTERM:
		debug("handleSignal: Received SIGTERM")
		run = false
	case syscall.SIGHUP:
		debug("handleSignal: Received SIGHUP")
		if config.Debug {
			fmt.Println("Disabling the debug output")
			config.Debug = false
		} else {
			fmt.Println("Enabling the debug output")
			config.Debug = true
		}
	default:
		debug("handleSignal: Received unhandled signal!")
	}
}

/*
	Prepare the configuration with hardcoded defaults and try to parse the configuration file
*/
func parseConfig() error {
	/*
		Pre-set the hardcoded default values
	*/
	config.HTTP.IP = "127.0.0.1"
	config.HTTP.Port = 8080
	config.SMTP.Host = "127.0.0.1"
	config.SMTP.Port = 25
	config.SMTP.Agent = fmt.Sprintf("http2smtp proxy %s", version)

	/*
		Parse the command line arguments
	*/
	var f string
	flag.StringVar(&f, "config", "/etc/http2smtp/config.json", "Path to the configuration file")
	flag.Parse()

	/*
		Try to read the configuration file
	*/
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	/*
		Try to parse the configuration file
	*/
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	/*
		Parse the environment
	*/
	if _, ok := os.LookupEnv("DEBUG"); ok {
		config.Debug = true
	}

	config.host, err = os.Hostname()
	if err != nil {
		return err
	}

	config.HTTP.listen = fmt.Sprintf("%s:%d", config.HTTP.IP, config.HTTP.Port)

	config.SMTP.server, err = gosmtp.NewServer(fmt.Sprintf("%s:%d", config.SMTP.Host, config.SMTP.Port), config.SMTP.Agent)
	if err != nil {
		return err
	}

	return nil
}
