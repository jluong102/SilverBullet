package main

import (
	"fmt"
	"os"
)

// Settings loaded in from config file
type Settings struct {
	Bullets []string `json:"bullets"`
	Log     string   `json:"log"`
}

// Make sure all needed settings are good
func (this Settings) VerifySettings() {
	fmt.Printf("Verifying settings\n")

	this.verifyBullets()
	this.verifyLog()
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

	// Check if log exists
	info, err := os.Stat(this.Log)

	if err == nil {
		// Make sure path is found
		if info.IsDir() {
			fmt.Printf("Verifyed log path %s\n", this.Log)
		} else {
			fmt.Printf("%s is not a directory\n", this.Log)
			os.Exit(INVALID_PATH_ERROR)
		}
	} else {
		// Directory not found, make one
		fmt.Printf("Creating directory %s\n", this.Log)

		if err = os.MkdirAll(this.Log, 0744); err != nil {
			fmt.Printf("Failed to create directory %s: %s\n", this.Log, err)
			os.Exit(FILE_CREATE_ERROR)
		}
	}
}
