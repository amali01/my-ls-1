package get

import (
	"os"
	"sort"
	"strings"
)

func MySort(files []File, flagsToUse Flags, ZeroArgs bool) []File {
	// Sort the files slice based on NotThere and CWD.
	var temp File
	var tempIndex int

	if ZeroArgs || flagsToUse.Flag_R {
		temp = files[0]
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].NotThere != files[j].NotThere {
			// Sort by NotThere first (false comes before true).
			return files[i].NotThere
		}
		if files[i].NotFolder != files[j].NotFolder {
			// Sort by NotFolder first (false comes before true).
			return files[i].NotFolder
		}
		if flagsToUse.Flag_t {

			fileInfo1, err1 := os.Lstat(files[i].CWD)
			fileInfo2, err2 := os.Lstat(files[j].CWD)

			// Handle errors, e.g., if there's a problem getting file info.
			if err1 != nil || err2 != nil {
				return false // You may want to handle errors differently.
			}

			// Compare files by ModTime.
			if flagsToUse.Flag_r {
				// If Flag_r is true, reverse the order.
				return fileInfo1.ModTime().Before(fileInfo2.ModTime())
			}

			return fileInfo1.ModTime().After(fileInfo2.ModTime())
		}

		if flagsToUse.Flag_r && i != 0 {
			// If Flag_r is true, reverse the order.
			return strings.ToLower(files[i].CWD) > strings.ToLower(files[j].CWD)
		}
		// Sort CWD as if they were lowercase.
		return strings.ToLower(files[i].CWD) < strings.ToLower(files[j].CWD)
	})

	if ZeroArgs || flagsToUse.Flag_R {
		// Find the index of the file that matches temp
		for i, f := range files {
			if f.CWD == temp.CWD {
				tempIndex = i
				break
			}
		}
	}

	if tempIndex != 0 {
		// Remove temp from its current position
		tempFile := files[tempIndex]
		files = append(files[:tempIndex], files[tempIndex+1:]...)

		// Insert temp at index 0
		files = append([]File{tempFile}, files...)
	}

	return files
}
