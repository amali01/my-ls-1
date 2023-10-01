package get

import (
	"fmt"
	"os"
	"strings"
)

func MyFlags(args []string, flagsToUse Flags) ([]string, Flags) {
	for i, arg := range args {
		if strings.HasPrefix(arg, "-") {

			flagsToUse = DetectFlag(arg, flagsToUse)
			args[i] = ""
		}
	}
	args = CleanedInputargs(args)
	return args, flagsToUse
}

// Detects if flags from user are right
func DetectFlag(flagToCheck string, flagsToUse Flags) Flags {
	dashFound := false
	for _, i := range flagToCheck {

		if i == 'R' && flagsToUse.Flag_R == false {
			flagsToUse.Flag_R = true
		} else if i == 'a' && flagsToUse.Flag_a == false {
			flagsToUse.Flag_a = true
		} else if i == 'r' && flagsToUse.Flag_r == false {
			flagsToUse.Flag_r = true
		} else if i == 'l' && flagsToUse.Flag_l == false {
			flagsToUse.Flag_l = true
		} else if i == 't' && flagsToUse.Flag_t == false {
			flagsToUse.Flag_t = true
		} else if i == '-' && !dashFound {
			dashFound = true
		} else {

			fmt.Printf("ERROR: %v is not a correct flag\n", i)
			os.Exit(0)
		}
	}

	return flagsToUse
}
