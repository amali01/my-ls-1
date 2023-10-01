package get

import (
	"sort"
	"strings"
)

func MySort(files []File, flagsToUse Flags, ZeroArgs bool, mainRoot string) []File {
	// Sort the files slice based on NotThere and CWD.
	var temp File
	var tempIndex int

	if ZeroArgs {
		temp = files[0]
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].NotThere != files[j].NotThere {
			// Sort by NotThere first (false comes before true).
			return files[i].NotThere
		}
		if flagsToUse.Flag_r && i != 0 {
			// If Flag_r is true, reverse the order.
			return strings.ToLower(files[i].CWD) > strings.ToLower(files[j].CWD)
		}
		// Sort CWD as if they were lowercase.
		return strings.ToLower(files[i].CWD) < strings.ToLower(files[j].CWD)
	})

	if ZeroArgs {
		// Find the index of the file that matches temp
		for i, f := range files {
			if f.CWD == temp.CWD {
				tempIndex = i
				break
			}
		}
	}

	if ZeroArgs && tempIndex != 0 {
		// Remove temp from its current position
		tempFile := files[tempIndex]
		files = append(files[:tempIndex], files[tempIndex+1:]...)

		// Insert temp at index 0
		files = append([]File{tempFile}, files...)
	}

	return files
}
