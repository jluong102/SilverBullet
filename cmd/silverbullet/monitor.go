package main

import (
	"fmt"
	"os"
)

type Monitor struct {
	Interval uint             `json:"interval"`
	Script   string           `json:"script"`
	Good     []int            `json:"good"`
	Bad      map[string][]int `json:"bad"`
}

func (this Monitor) VerifyMonitor() {
	this.verifyScript()
	this.verifyGood()
}

func (this Monitor) verifyScript() {
	fmt.Printf("Verifying monitor script %s\n", this.Script)
	VerifyScript(this.Script)
}

// Make sure that at least one good exit code is provided
func (this Monitor) verifyGood() {
	if len(this.Good) < 1 {
		fmt.Printf("Must provide \"good\" exit codes in \"monitor\"\n")
		os.Exit(MISSING_SETTING_ERROR)
	}
}
