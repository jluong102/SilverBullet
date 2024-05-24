package main

import (
	"fmt"
)

type Remedy struct {
	Script   string `yaml:"script"`
	Try      uint   `yaml:"try"`
	Interval uint   `yaml:"interval"`
}

func (this Remedy) VerifyRemedy() {
	this.verifyScript()
}

func (this Remedy) verifyScript() {
	fmt.Printf("\tVerifying remedy script %s\n", this.Script)
	VerifyScript(this.Script)
}
