package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
	"time"
)

type Bullet struct {
	Name    string
	Monitor *Monitor           `yaml:"monitor"`
	Remedy  map[string]*Remedy `yaml:"remedy"`
}

// Bullet stuff
func (this Bullet) VerifyBullet() {
	this.Monitor.VerifyMonitor()
	this.VerifyRemedy()
}

/*
 * We want to make sure that each remedy
 * that is declared inside monitoring config
 * is valid. This needs to be checked at the
 * "bullet" level so we can read from both
 * "monitor" and "remedy". This should be called
 * after the verifaction is done on "monitor".
 */
func (this Bullet) VerifyRemedy() {
	for i, _ := range this.Monitor.Bad {
		if _, found := this.Remedy[i]; found {
			fmt.Printf("Verifying remedy %s\n", i)
			this.Remedy[i].VerifyRemedy()
		} else {
			fmt.Printf("%s not defined in \"remedy\"\n", i)
			os.Exit(MISSING_REMEDY_ERROR)
		}
	}
}

func (this Bullet) StartScan(wg *sync.WaitGroup) {
	fmt.Printf("Starting bullet %s", this.Name)
	defer wg.Done() // This shouldn't be needed

	for {
		if this.isOOR() {
			time.Sleep(time.Minute)
			continue
		}

		remedy := this.Monitor.RunMonitor()

		// Run remedy if defined
		if remedy != "" {
			this.Remedy[remedy].RunRemedy()
		}
	}
}

func (this Bullet) isOOR() bool {
	return false
}

// Non object stuff
func LoadBullet(filename string) *Bullet {
	fmt.Printf("\tLoading bullet %s\n", filename)

	bullet := new(Bullet)
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Unable to read from file %s: %s\n", filename, err)
		os.Exit(FILE_READ_ERROR)
	}

	if err = yaml.Unmarshal(content, bullet); err != nil {
		fmt.Printf("Trouble parsing: %s\n", err)
		os.Exit(YAML_PARSE_ERROR)
	}

	bullet.Name = filename

	return bullet
}
