package main

import (
	"fmt"
	"myLS/get"
	"os"
	"sort"
	"strings"
)

type file struct {
	CWD      string
	Names    []string
	Types    []string
	Color    []string
	NotThere bool
}

func main() {
	args := os.Args[1:]
	var files []file
	var err error

	if len(os.Args) == 1 {

		// Initialize the files slice with a file struct.
		files = append(files, file{})

		files[0].CWD, err = os.Getwd()
		if err != nil {
			fmt.Println("Error1:", err)
			files[0].NotThere = true
			// return
		}

		files[0].Names, err = get.GetInfo(files[0].CWD)
		if err != nil {
			fmt.Println("Error4:", err)
			files[0].NotThere = true
			// return
		} else {
			files[0].Types, files[0].Color, err = get.GetFileType(files[0].Names)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		// fmt.Println("1:", files[0].Names)///////////////////////////////////////////////////////////////////////////////////////////
		// fmt.Println("2:", files[0].Types)

		get.DisplayInfo(files[0].Names, files[0].Color)

	} else {
		for i, arg := range args {
			files = append(files, file{})
			files[i].CWD = arg
			files[i].Names, err = get.GetInfo(arg)
			if err != nil {
				fmt.Println("Error5:", err)
				files[i].NotThere = true
			}
		}

		// fmt.Println("befor", files) //////////////////////////////////////////

		// Sort the files slice based on NotThere and CWD.
		sort.Slice(files, func(i, j int) bool {
			if files[i].NotThere != files[j].NotThere {
				// Sort by NotThere first (false comes before true).
				return files[i].NotThere
			}

			// Sort CWD as if they were lowercase.
			return strings.ToLower(files[i].CWD) < strings.ToLower(files[j].CWD)
		})

		// fmt.Println("after", files) //////////////////////////////////////////

		for i, f := range files {
			if f.NotThere {
				fmt.Printf("ls: cannot access '%v': No such file or directory\n", f.CWD)
			} else {
				if len(files) > 1 && i != len(files)-1 {
					fmt.Printf("%v:\n", f.CWD)
					get.DisplayInfo(f.Names, f.Color)
					fmt.Println()
				} else if len(files) > 1 && i == len(files)-1 {
					fmt.Printf("%v:\n", f.CWD)
					get.DisplayInfo(f.Names, f.Color)
				} else {
					get.DisplayInfo(f.Names, f.Color)

				}
			}
		}

	}

}

// func main() {
// 	args := os.Args
// 	if len(args) < 1 {
// 		fmt.Println("Error\nUsage: go run . [OPTIONS] [FILE|DIR]")
// 		os.Exit(0)
// 	}

// 	// Get the current working directory.
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	// fmt.Println("Current working directory:", cwd)

// 	// // Open the current working directory.
// 	dir, err := os.Open(cwd)
// 	if err != nil {
// 		fmt.Println("Error1:", err)
// 		return
// 	}
// 	defer dir.Close()

// 	// Read the contents of the directory.
// 	entries, err := dir.ReadDir(0) // Read all entries in the directory.
// 	if err != nil {
// 		fmt.Println("Error2:", err)
// 		return
// 	}

// 	// Create a slice to store the names of both files and directories.
// 	var names []string

// 	// Collect both file and directory names (excluding hidden files).
// 	for _, entry := range entries {
// 		name := entry.Name()
// 		if !strings.HasPrefix(name, ".") { // Exclude hidden files.
// 			names = append(names, name)
// 		}
// 	}

// 	// Sort the names as if they were lowercase but DisplayInfo the actual values.
// 	sort.SliceStable(names, func(i, j int) bool {
// 		return strings.ToLower(names[i]) < strings.ToLower(names[j])
// 	})

// 	// Print the original case names.
// 	for i, name := range names {
// 		if i == 0 {
// 			fmt.Printf("%v ", name)
// 		} else if i == len(names)-1 {
// 			fmt.Printf(" %v", name)
// 		} else {
// 			fmt.Printf(" %v ", name)
// 		}

// 	}
// 	defer fmt.Println()

// }
