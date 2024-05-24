package main

import (
	"fmt"
	"os"
	"sync"
	"gopkg.in/yaml.v3"
)

type Bullet struct {
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
	defer wg.Done() // This shouldn't be needed

	for {
		_ = this.Monitor.RunMonitor()
	}
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

	return bullet
}
