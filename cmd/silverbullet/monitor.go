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
	fmt.Printf("Verifying monitor\n")
	this.verifyScript()
	this.verifyGood()
	this.verifyBad()
}

func (this Monitor) verifyScript() {
	fmt.Printf("\tVerifying monitor script %s\n", this.Script)
	VerifyScript(this.Script)
}

// Make sure that at least one good exit code is provided
func (this Monitor) verifyGood() {
	if len(this.Good) < 1 {
		fmt.Printf("\tMust provide \"good\" exit codes in \"monitor\"\n")
		os.Exit(MISSING_SETTING_ERROR)
	}
}

func (this Monitor) verifyBad() {
	for i, j := range this.Bad {
		if len(j) < 1 {
			fmt.Printf("\tMust provide at least one exitcode for %s\n", i)
			os.Exit(MISSING_SETTING_ERROR)
		}
	}
}
