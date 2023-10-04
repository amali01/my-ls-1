package main

import (
	"fmt"
	"myLS/get"
	"os"
)

func main() {
	// Parse command-line arguments and flags
	args := os.Args[1:]
	var files []get.File
	var flagsToUse get.Flags
	var ZeroArgs bool // Indicates if running the program with zero arguments or only with flags
	SubCount := 0     // Subfolder counter
	var err error

	// Parse command-line arguments and flags
	args, flagsToUse = get.MyFlags(args, flagsToUse) // Extract and categorize the used flags

	// Store the original length of input arguments without flags
	OgArgsLen := len(args)

	// Initialize the root directory information
	var mainRoot get.File
	mainRoot.CWD, _ = get.MyPath()

	// If no arguments are provided, set ZeroArgs to true and use the current working directory as an argument
	if len(args) == 0 {
		ZeroArgs = true
		args = append(args, mainRoot.CWD)
	}

	// Loop through the provided arguments
	for i := 0; i < len(args); i++ {
		arg := args[i]
		files = append(files, get.File{})
		files[i].CWD = arg

		// Retrieve file information and categorize it based on flags
		files[i], err = get.GetInfo(files[i], flagsToUse, i)

		if err != nil {
			files[i].NotThere = true
		} else {
			if !files[i].NotFolder || len(files[i].Names) > 1 { // Enter if it's a folder and not a symbolic link
				files[i], args, err = get.GetFileType(files[i], args, flagsToUse, mainRoot.CWD, i)
				if err != nil {
					fmt.Println("Error111:", err)
					return
				}
			} else {
				files[i].FileName = arg
			}
		}

		// Check if the number of arguments has changed (indicating a subfolder)
		if OgArgsLen != len(args) {
			OgArgsLen = len(args)
			SubCount++
		}
	}

	// Sort the files and print the output
	files = get.MySort(files, flagsToUse, ZeroArgs)
	get.MyPrint(files, flagsToUse, ZeroArgs)
}
