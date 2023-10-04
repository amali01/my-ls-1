package get

import (
	"fmt"
	"os"
	"strings"
)

// MyFlags parses command-line arguments and extracts the ls command flags.
func MyFlags(args []string, flagsToUse Flags) ([]string, Flags) {
	var dashFound bool
	for i, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flagsToUse, dashFound = DetectFlag(arg, flagsToUse)
			if dashFound {
				continue
			}
			args[i] = ""
		}
	}
	args = CleanedInputargs(args)
	return args, flagsToUse
}

// DetectFlag checks if a flag is valid and sets it in the Flags struct.
func DetectFlag(flagToCheck string, flagsToUse Flags) (Flags, bool) {
	dashFound := false
	if flagToCheck == "-" && !dashFound {
		return flagsToUse, true
	} else if flagToCheck == "--help" {
		fmt.Println("Available Flags\n-l\n-R\n-a\n-r\n-t\n")
		os.Exit(0)
	}
	for _, i := range flagToCheck {
		if i == 'R' && !flagsToUse.Flag_R {
			flagsToUse.Flag_R = true
		} else if i == 'a' && !flagsToUse.Flag_a {
			flagsToUse.Flag_a = true
		} else if i == 'r' && !flagsToUse.Flag_r {
			flagsToUse.Flag_r = true
		} else if i == 'l' && !flagsToUse.Flag_l {
			flagsToUse.Flag_l = true
		} else if i == 't' && !flagsToUse.Flag_t {
			flagsToUse.Flag_t = true
		} else if i == '-' && !dashFound {
			dashFound = true
		} else {
			fmt.Printf("MyLs: invalid option -- '%v'\n", string(i))
			fmt.Println("Try 'go run . --help' for more information.")
			os.Exit(0)
		}
	}
	return flagsToUse, false
}

// CalculateBlocks calculates the total disk usage of listed files.
func CalculateBlocks(files []File) int {
	var sum int
	for _, f := range files {
		sum = sum + f.TotalSize
	}
	return sum
}
