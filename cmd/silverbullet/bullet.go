package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Bullet struct {
	Monitor *Monitor `json:"monitor"`
}

// Bullet stuff
func (this Bullet) VerifyBullet() {
	this.Monitor.VerifyMonitor()
}

// Non object stuff
func LoadBullet(filename string) *Bullet {
	fmt.Printf("Loading bullet from %s\n", filename)

	bullet := new(Bullet)
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Unable to read from file %s: %s\n", filename, err)
		os.Exit(FILE_READ_ERROR)
	}

	if err = json.Unmarshal(content, bullet); err != nil {
		fmt.Printf("Trouble parsing: %s\n", err)
		os.Exit(JSON_PARSE_ERROR)
	}

	return bullet
}
