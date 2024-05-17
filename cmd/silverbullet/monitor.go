package main

import (
	"fmt"
	"os"
)

type Monitor struct {
	Interval uint   `json:"interval"`
	Script   string `json:"script"`
}

func (this Monitor) VerifyMonitor() {
	this.verifyScript()
}

func (this Monitor) verifyScript() {
	info, err := os.Stat(this.Script)

	if err != nil {
		fmt.Printf("Unable to find file %s: %s\n", this.Script, err)
		os.Exit(SCRIPT_FILE_ERROR)
	} else if info.IsDir() {
		fmt.Printf("Script %s cannont be a directory\n", this.Script)
		os.Exit(SCRIPT_FILE_ERROR)
	}
}
