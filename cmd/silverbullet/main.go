package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"
)

// Set at compile time
var VERSION string = "N/A"
var BUILD_DATE string = "N/A"

// For holding args from cmdline
type cmdline struct {
	config  string
	version bool
}

// Load in args passed from cmdline
func loadArgs(cmdArgs *cmdline) {
	flag.StringVar(&cmdArgs.config, "c", "/etc/silverbullet/config.json", "Setup config to use")
	flag.BoolVar(&cmdArgs.version, "version", false, "Print version")

	flag.Parse()

	if cmdArgs.version {
		printVersion()
	}
}

// Print version an exit
func printVersion() {
	fmt.Printf("~ Silver Bullet ~\n")
	fmt.Printf("Version: %s\n", VERSION)
	fmt.Printf("Date: %s\n", BUILD_DATE)

	os.Exit(NO_ERROR)
}

// Make sure valid cmdline args are passed
func checkArgs(cmdArgs *cmdline) {

}

// Read in content from config file
func loadSettings(configFile string) *Settings {
	fmt.Printf("Loading settings from %s\n", configFile)

	var settings *Settings = new(Settings)
	content, err := os.ReadFile(configFile)

	if err != nil {
		fmt.Printf("Unable to read from file %s: %s\n", configFile, err)
		os.Exit(FILE_READ_ERROR)
	}

	if err = json.Unmarshal(content, settings); err != nil {
		fmt.Printf("Trouble parsing: %s\n", err)
		os.Exit(JSON_PARSE_ERROR)
	}

	return settings
}

func initScans(bullets []Bullet) {
	fmt.Printf("Starting ups scans\n")
	var wg sync.WaitGroup

	for _, i := range bullets {
		wg.Add(1)
		go i.StartScan(&wg)
	}

	wg.Wait()
}

func main() {
	cmdArgs := new(cmdline)

	loadArgs(cmdArgs)
	checkArgs(cmdArgs)

	settings := loadSettings(cmdArgs.config)
	settings.VerifySettings()

	bullets := settings.GetBullets()
	initScans(bullets)
}
