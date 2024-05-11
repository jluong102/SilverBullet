package main

import (
	"fmt"
	"os"
)

// Settings loaded in from config file
type Settings struct {
	Bullets []string `json:"bullets"`
}

// Make sure all needed settings are good
func (this Settings) VerifySettings() {
	fmt.Printf("Verifying settings\n")

	this.verifyBullets()
}

// Create bullet objects from config file and return them
func (this Settings) GetBullets() []Bullet {
	var bullets []Bullet

	for _, i := range this.Bullets {
		b := LoadBullet(i)
	}

	return bullets
}

// Make sure at least one bullet is provided
func (this Settings) verifyBullets() {
	if len(this.Bullets) < 1 {
		fmt.Printf("Must provide \"bullets\" in config\n")
		os.Exit(MISSING_SETTING_ERROR)
	}
}
