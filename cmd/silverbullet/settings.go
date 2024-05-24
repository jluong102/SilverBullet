package main

import (
	"fmt"
	"os"
)

// Settings loaded in from config file
type Settings struct {
	Bullets []string `yaml:"bullets"`
	Log     string   `yaml:"log"`
	OOR     string   `yaml:"oor,omitempty"`
}

// Make sure all needed settings are good
func (this Settings) VerifySettings() {
	fmt.Printf("Verifying settings\n")

	this.verifyBullets()
	this.verifyLog()
	this.verifyOOR()
}

// Create bullet objects from config file and return them
func (this Settings) GetBullets() []Bullet {
	fmt.Printf("Loading bullets\n")

	var bullets []Bullet

	for _, i := range this.Bullets {
		b := LoadBullet(i)
		b.VerifyBullet()

		bullets = append(bullets, *b)
	}

	return bullets
}

// Make sure at least one bullet is provided
func (this Settings) verifyBullets() {
	if len(this.Bullets) < 1 {
		fmt.Printf("\tMust provide \"bullets\" in config\n")
		os.Exit(MISSING_SETTING_ERROR)
	}
}

func (this Settings) verifyLog() {
	// Log is optional
	if this.Log == "" {
		fmt.Printf("\tNo \"log\" set in config\n")
	}

	// Check if log exists or create it
	VerifyDirPath(this.Log)
}

func (this Settings) verifyOOR() {
	if this.OOR == "" { // Default to /etc/silverbullet/oor
		this.OOR = "/etc/silverbullet/oor"
		fmt.Printf("\tNo \"oor\" path set. Using default %s\n", this.OOR)
	}

	// Check if path exists or create it
	VerifyDirPath(this.OOR)
}
