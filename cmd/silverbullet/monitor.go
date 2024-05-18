package main

import (
	"fmt"
	"os"
)

type Monitor struct {
	Interval uint   `json:"interval"`
	Script   string `json:"script"`
	Good     []int  `json:"good"`
}

func (this Monitor) VerifyMonitor() {
	this.verifyScript()
	this.verifyGood()
}

func (this Monitor) verifyScript() {
	fmt.Printf("Verifying monitor script %s\n", this.Script)

	info, err := os.Stat(this.Script)

	if err != nil {
		fmt.Printf("Unable to find file %s: %s\n", this.Script, err)
		os.Exit(SCRIPT_FILE_ERROR)
	} else if info.IsDir() {
		fmt.Printf("Script %s cannot be a directory\n", this.Script)
		os.Exit(SCRIPT_FILE_ERROR)
	}
}

// Make sure that at least one good exit code is provided
func (this Monitor) verifyGood() {
	if len(this.Good) < 1 {
		fmt.Printf("Must provide \"good\" exit codes in \"monitor\"\n")
		os.Exit(MISSING_SETTING_ERROR)
	}
}
