package get

import (
	"fmt"
)

func MyPrint(files []File, flagsToUse Flags, ZeroArgs bool) {
	for i, f := range files {
		if flagsToUse.Flag_l && ZeroArgs && i == 0 && len(files) == 1 {
			fmt.Println("total:", CalculateBlocks(files))
		}
		if ZeroArgs {
			if flagsToUse.Flag_R && i == 0 {
				f.FileName = "."
			} else {
				f.FileName = "." + f.FileName
			}
		}

		if f.NotThere && !files[i].NotFolder { //file or directory does not exist.
			fmt.Printf("ls: cannot access '%v': No such file or directory\n", f.CWD)

		} else if files[i].NotFolder { //if it's not a directory or folder
			if flagsToUse.Flag_l {
				if f.SoftLinks[i] == "" { //if not a softlink
					fmt.Println(f.Permission[i], f.Hardlinks[i], f.Group[i], f.Owner[i], f.AlingSize[i], f.Month[i], f.Day[i], f.ModificationTime[i], f.CWD)
					// fmt.Println(f.Permission[0], f.Hardlinks[0], f.Group[0], f.Owner[0], f.Size[0], f.Month[0], f.Day[0], f.ModificationTime[0], f.CWD)
				} else {
					fmt.Println(f.Permission[i], f.Hardlinks[i], f.Group[i], f.Owner[i], f.AlingSize[i], f.Month[i], f.Day[i], f.ModificationTime[i], f.FileColor[i], f.CWD, "\033[0m", "->", f.SoftLinks[i])
				}
			} else {
				fmt.Println(f.FileName) //without flag -l
			}

		} else {
			if len(files) > 1 && i != len(files)-1 { // if multiple folders but not the last one
				fmt.Printf("%v:\n", f.FileName)
				if len(files) > 1 && flagsToUse.Flag_l {
					fmt.Println("total:", f.TotalSize)
				}
				if flagsToUse.Flag_l {
					FullInfo(f)
				} else {
					DisplayInfo(f.Names, f.FileColor)
				}

				fmt.Println() // the difference

			} else if len(files) > 1 && i == len(files)-1 { // at last one of multiple folders
				fmt.Printf("%v:\n", f.FileName)
				if flagsToUse.Flag_l {
					fmt.Println("total:", f.TotalSize)
				}
				if flagsToUse.Flag_l {
					FullInfo(f)
				} else {
					DisplayInfo(f.Names, f.FileColor)
				}

			} else {
				if flagsToUse.Flag_R {
					fmt.Printf("%v:\n", f.FileName)
				}

				if flagsToUse.Flag_l && !ZeroArgs {
					fmt.Println("total:", f.TotalSize)
				}

				if flagsToUse.Flag_l { // if flag -l with zero args
					FullInfo(f)
				} else {
					DisplayInfo(f.Names, f.FileColor) //if zero args with no flags or (flags -a -t -r)
				}

			}
		}
	}
}

func DisplayInfo(names []string, color []string) {
	// Print the original case names. if flag l is not used.
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

func FullInfo(file File) {
	// Print the Full Informations. when flag l is used.
	for i, name := range file.Names {
		if file.FileColor[i] == "reset" {
			if file.SoftLinks[i] == "" {
				fmt.Println(file.Permission[i], file.Hardlinks[i], file.Group[i], file.Owner[i], file.AlingSize[i], file.Month[i], file.Day[i], file.ModificationTime[i], name)
			} else {
				fmt.Println(file.Permission[i], file.Hardlinks[i], file.Group[i], file.Owner[i], file.AlingSize[i], file.Month[i], file.Day[i], file.ModificationTime[i], name, "->", file.SoftLinks[i])
			}
		} else {
			if file.SoftLinks[i] == "" {
				fmt.Println(file.Permission[i], file.Hardlinks[i], file.Group[i], file.Owner[i], file.AlingSize[i], file.Month[i], file.Day[i], file.ModificationTime[i], file.FileColor[i], name, "\033[0m")
			} else {
				fmt.Println(file.Permission[i], file.Hardlinks[i], file.Group[i], file.Owner[i], file.AlingSize[i], file.Month[i], file.Day[i], file.ModificationTime[i], file.FileColor[i], name, "\033[0m", "->", file.SoftLinks[i])
			}
		}
	}

}
