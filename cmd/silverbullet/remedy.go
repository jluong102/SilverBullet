package main

import (
	"fmt"
)

type Remedy struct {
	Script   string `json:"script"`
	Try      uint   `json:"try"`
	Interval uint   `json:"interval"`
}

func (this Remedy) VerifyRemedy() {
	this.verifyScript()
}

func (this Remedy) verifyScript() {
	fmt.Printf("Verifying remedy script %s", this.Script)
	VerifyScript(this.Script)
}
