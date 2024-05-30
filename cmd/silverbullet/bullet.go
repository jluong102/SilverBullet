package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
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

/*
 * This is the main function of each bullet.
 * This will start by checking to see if the bullet
 * is marked OOR. If the bullet is not marked OOR,
 * run the monitoring script/bin until it fails.
 * The remedy name is returned and then run.
 * If this fails more then the set "try" amount marked
 * mark the bullet OOR.
 */
func (this Bullet) StartScan(settings *Settings, wg *sync.WaitGroup) {
	fmt.Printf("Starting bullet %s\n", this.Name)
	defer wg.Done() // This shouldn't be needed
	var attempt uint = 0

	for {
		if this.isOOR(settings.OOR) {
			fmt.Printf("%s is currently marked OOR\n", this.Name)

			time.Sleep(time.Minute)
			continue
		}

		remedy := this.Monitor.RunMonitor(&attempt)

		// Run remedy if defined
		if remedy != "" {
			this.Remedy[remedy].RunRemedy()
		} else {
			fmt.Printf("Unknown exitcode\n")
		}

		attempt++

		if attempt >= this.Remedy[remedy].Try && this.Remedy[remedy].Try != 0 {
			this.markOOR(settings.OOR)
		}
	}
}

func (this Bullet) isOOR(dirPath string) bool {
	content, err := os.ReadDir(dirPath)

	if err != nil {
		fmt.Printf("Failed to read from directory: %s\n", err)
		os.Exit(INVALID_PATH_ERROR)
	}

	for _, info := range content {
		if info.Name() == this.Name {
			return true
		}
	}

	return false
}

// Create a file an the OOR directory of the bullet name
func (this Bullet) markOOR(dirOOR string) {
	fmt.Printf("Marking %s OOR\n", this.Name)
	oorFile := fmt.Sprintf("%s/%s", dirOOR, this.Name)

	// File should not be created at this point
	if _, err := os.Stat(oorFile); err != nil {
		fmt.Printf("OOR file already found\n")
		return
	}

	oor, err := os.Create(oorFile)

	if err != nil {
		fmt.Printf("Unable to mark %s OOR: %s\n", this.Name, err)
	}

	oor.WriteString("Failed to remedy")
	oor.Close()
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

	// Set the name of the bullet to the filename with no extension
	tmp := strings.Split(filename, "/")
	bullet.Name = tmp[len(tmp)-1]
	bullet.Name = bullet.Name[:len(bullet.Name)-5]

	return bullet
}
