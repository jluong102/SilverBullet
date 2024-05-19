package main

import (
	"fmt"
	"os"
)

// Settings loaded in from config file
type Settings struct {
	Bullets []string `json:"bullets"`
	Log     string   `json:"log"`
	OOR     string   `json:"oor"`
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
		fmt.Printf("Must provide \"bullets\" in config\n")
		os.Exit(MISSING_SETTING_ERROR)
	}
}

func (this Settings) verifyLog() {
	// Log is optional
	if this.Log == "" {
		fmt.Printf("No \"log\" set in config\n")
	}

	// Check if log exists or create it
	VerifyDirPath(this.Log)
}

func (this Settings) verifyOOR() {
	if this.OOR == "" { // Default to /var/silverbullet/oor
		this.OOR = "/var/silverbullet/oor"
		fmt.Printf("No \"oor\" path set. Using default %s\n", this.OOR)
	}

	// Check if path exists or create it
	VerifyDirPath(this.OOR)
}
