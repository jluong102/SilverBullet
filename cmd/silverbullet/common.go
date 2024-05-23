package main

import (
	"fmt"
	"os"
)

/*
 * Use this to make sure that the file is found
 * and that file is not a directory.
 * Try to make sure you are able to execute the
 * provded file as well
 */
func VerifyScript(script string) {
	info, err := os.Stat(script)

	if err != nil {
		fmt.Printf("\tUnable to find file %s: %s\n", script, err)
		os.Exit(SCRIPT_FILE_ERROR)
	} else if info.IsDir() {
		fmt.Printf("\tScript %s cannot be a directory\n", script)
		os.Exit(SCRIPT_FILE_ERROR)
	}
}

/*
 * Use this to make sure either the provided dir
 * exists or that we are able to create the path.
 * Exit program on fail.
 */
func VerifyDirPath(dirPath string) {
	info, err := os.Stat(dirPath)

	if err == nil {
		// Make sure path is found
		if info.IsDir() {
			fmt.Printf("\tVerifyed path %s\n", dirPath)
		} else {
			fmt.Printf("\t%s is not a directory\n", dirPath)
			os.Exit(INVALID_PATH_ERROR)
		}
	} else {
		// Directory not found, make one
		fmt.Printf("\tCreating directory %s\n", dirPath)

		if err = os.MkdirAll(dirPath, 0744); err != nil {
			fmt.Printf("\tFailed to create directory %s: %s\n", dirPath, err)
			os.Exit(FILE_CREATE_ERROR)
		}
	}
}
