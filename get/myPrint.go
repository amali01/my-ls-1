package get

import "fmt"

func MyPrint(files []File, flagsToUse Flags, ZeroArgs bool) {
	for i, f := range files {
		if ZeroArgs {
			if flagsToUse.Flag_R && i == 0 {
				f.CWD = "."
			} else {
				f.CWD = "./" + f.CWD
			}
		}
		if f.NotThere {
			fmt.Printf("ls: cannot access '%v': No such file or directory\n", f.CWD)
		} else {
			if len(files) > 1 && i != len(files)-1 {
				fmt.Printf("%v:\n", f.CWD)
				DisplayInfo(f.Names, f.FileColor)
				fmt.Println()
			} else if len(files) > 1 && i == len(files)-1 {
				fmt.Printf("%v:\n", f.CWD)
				DisplayInfo(f.Names, f.FileColor)
			} else {
				DisplayInfo(f.Names, f.FileColor)

			}
		}
	}
}

func DisplayInfo(names []string, color []string) {

	// Print the original case names.
	for i, name := range names {
		if color[i] == "" {
			if i == 0 {
				fmt.Printf("%v ", name)
			} else if i == len(names)-1 {
				fmt.Printf(" %v", name)
			} else {
				fmt.Printf(" %v ", name)
			}
		} else {
			if i == 0 {
				fmt.Printf("%v%v%v ", color[i], name, "\033[0m") // first one
			} else if i == len(names)-1 {
				fmt.Printf(" %v%v%v", color[i], name, "\033[0m") // last one
			} else {
				fmt.Printf(" %v%v%v ", color[i], name, "\033[0m") // the rest
			}
		}

	}
	fmt.Println()

}
