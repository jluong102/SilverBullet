// ///////////////////////////////////////////////////////
//
//	____  _ _                ____        _ _      _    //
//
// / ___|(_) |_   _____ _ __| __ ) _   _| | | ___| |_  //
// \___ \| | \ \ / / _ \ '__|  _ \| | | | | |/ _ \ __| //
//
//	___) | | |\ V /  __/ |  | |_) | |_| | | |  __/ |_  //
//
// |____/|_|_| \_/ \___|_|  |____/ \__,_|_|_|\___|\__| //
// ///////////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// For holding args from cmdline
type cmdline struct {
	config string
}

// Load in args passed from cmdline
func loadArgs(cmdArgs *cmdline) {
	flag.StringVar(&cmdArgs.config, "config", "./config.json", "Setup config to use")

	flag.Parse()
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

func main() {
	cmdArgs := new(cmdline)

	loadArgs(cmdArgs)
	checkArgs(cmdArgs)

	settings := loadSettings(cmdArgs.config)
	settings.VerifySettings()

	_ = settings.GetBullets()
}
