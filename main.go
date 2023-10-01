package main

import (
	"fmt"
	"myLS/get"
	"os"
)

func main() {
	args := os.Args[1:]
	var files []get.File
	var flagsToUse get.Flags
	var ZeroArgs bool // if running the program with  zero Argmintes or only with flags
	SupCount := 0     // supFolder counter
	var err error

	var mainRoot string
	args, flagsToUse = get.MyFlags(args, flagsToUse) // didacting the used flags
	OgArgsLen := len(args)                           // orgenal lenght of input without flags
	temp := OgArgsLen

	if len(args) == 0 {
		ZeroArgs = true
		arg, _ := get.MyPath()
		mainRoot = arg
		args = append(args, arg)
	}
	for i := 0; i < len(args); i++ {
		arg := args[i]
		files = append(files, get.File{})
		if ZeroArgs {
			if i == 0 {
				files[i].CWD = arg

			} else {

				if SupCount > 0 && OgArgsLen < i+1 {
					files[i].CWD = get.TrimPath(arg, true)
				} else {
					files[i].CWD = get.TrimPath(arg, false)
				}
			}
		} else {
			files[i].CWD = arg
		}

		files[i].Names, files[i].HiddenNames, err = get.GetInfo(arg, flagsToUse)
		if err != nil {
			files[i].NotThere = true
		} else {
			files[i].Types, files[i].FileColor, args, err = get.GetFileType(files[i].CWD, files[i].Names, args, flagsToUse, SupCount)
			if err != nil {
				fmt.Println("Error111:", err)
				return
			}
		}
		if ZeroArgs {
			if temp != len(args) {
				temp = len(args)
				if SupCount == 0 {
					OgArgsLen = len(args)
				}
				SupCount++
			}
		}
	}

	files = get.MySort(files, flagsToUse, ZeroArgs, mainRoot)

	get.MyPrint(files, flagsToUse, ZeroArgs)

}

// func main() {
// 	args := os.Args[1:]
// 	var files []get.File
// 	var err error
// 	var flagsToUse get.Flags
// 	var ZeroArgs bool
// 	SupCount := 0

// 	// var IsSubFolder bool
// 	args, flagsToUse = get.MyFlags(args, flagsToUse)
// 	OgArgsLen := len(args) // orgenal lenght of input without flags
// 	temp := OgArgsLen
// 	if len(args) < 1 {
// 		ZeroArgs = true
// 		if len(args) == 0 {
// 			arg, _ := get.MyPath()
// 			args = append(args, arg)
// 		}
// 		for i := 0; i < len(args); i++ {
// 			arg := args[i]
// 			files = append(files, get.File{})

// 			if i == 0 {
// 				files[i].CWD = arg

// 			} else {

// 				if SupCount > 0 && OgArgsLen < i+1 {
// 					files[i].CWD = get.TrimPath(arg, true)
// 				} else {
// 					files[i].CWD = get.TrimPath(arg, false)
// 				}
// 			}

// 			files[i].Names, files[i].HiddenNames, err = get.GetInfo(arg,flagsToUse)
// 			if err != nil {
// 				files[i].NotThere = true
// 			} else {
// 				files[i].Types, files[i].FileColor, args, err = get.GetFileType(files[i].CWD, files[i].Names, args, flagsToUse)
// 				if err != nil {
// 					fmt.Println("Error111:", err)
// 					return
// 				}
// 			}
// 			if temp != len(args) {
// 				temp = len(args)
// 				if SupCount == 0 {
// 					OgArgsLen = len(args)
// 				}
// 				SupCount++
// 			}
// 		}
// 		fmt.Println("count", SupCount)
// 		files = get.MySort(files, flagsToUse)

// 		get.MyPrint(files, flagsToUse, ZeroArgs)

// 	} else {

// 		for i := 0; i < len(args); i++ {
// 			arg := args[i]
// 			files = append(files, get.File{})
// 			files[i].CWD = arg
// 			files[i].Names, files[i].HiddenNames, err = get.GetInfo(arg)
// 			if err != nil {
// 				files[i].NotThere = true
// 			} else {
// 				files[i].Types, files[i].FileColor, args, err = get.GetFileType(files[i].CWD, files[i].Names, args, flagsToUse)
// 				if err != nil {
// 					fmt.Println("Error222:", err)
// 					return
// 				}
// 			}
// 		}

// 		// fmt.Println("befor", files) //////////////////////////////////////////

// 		files = get.MySort(files, flagsToUse)
// 		// // Sort the files slice based on NotThere and CWD.
// 		// sort.Slice(files, func(i, j int) bool {
// 		// 	if files[i].NotThere != files[j].NotThere {
// 		// 		// Sort by NotThere first (false comes before true).
// 		// 		return files[i].NotThere
// 		// 	}

// 		// 	// Sort CWD as if they were lowercase.
// 		// 	return strings.ToLower(files[i].CWD) < strings.ToLower(files[j].CWD)
// 		// })

// 		// fmt.Println("after", files) //////////////////////////////////////////

// 		get.MyPrint(files, flagsToUse, ZeroArgs)
// 		// for i, f := range files {
// 		// 	if f.NotThere {
// 		// 		fmt.Printf("ls: cannot access '%v': No such file or directory\n", f.CWD)
// 		// 	} else {
// 		// 		if len(files) > 1 && i != len(files)-1 {
// 		// 			fmt.Printf("%v:\n", f.CWD)
// 		// 			get.DisplayInfo(f.Names, f.FileColor)
// 		// 			fmt.Println()
// 		// 		} else if len(files) > 1 && i == len(files)-1 {
// 		// 			fmt.Printf("%v:\n", f.CWD)
// 		// 			get.DisplayInfo(f.Names, f.FileColor)
// 		// 		} else {
// 		// 			get.DisplayInfo(f.Names, f.FileColor)

// 		// 		}
// 		// 	}
// 		// }

// 	}

// }
