package main

import (
	"fmt"
	"time"
)

type Remedy struct {
	Script   string `yaml:"script"`
	Try      uint   `yaml:"try"`
	Interval uint   `yaml:"interval"`
}

func (this Remedy) RunRemedy() {
	RunScript(this.Script)

	// Wait before continuing
	time.Sleep(time.Second * time.Duration(this.Interval))
}

func (this Remedy) VerifyRemedy() {
	this.verifyScript()
}

func (this Remedy) verifyScript() {
	fmt.Printf("\tVerifying remedy script %s\n", this.Script)
	VerifyScript(this.Script)
}
