package get

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func GetInfo(cwd string, flagsToUse Flags) ([]string, []string, error) {
	// Open the current working directory.
	dir, err := os.Open(cwd)
	if err != nil {
		// fmt.Println("Error2:", err)
		return nil, nil, err
	}
	defer dir.Close()

	// Read the contents of the directory.
	// entries, err := dir.ReadDir(0) // Read all entries in the directory.
	entries, err := dir.Readdirnames(0) // Read all entries in the directory.
	if err != nil {
		fmt.Println("Error3:", err)
		return nil, nil, err
	}

	// Create a slice to store the names of both files and directories.
	var names []string
	var hiddenNames []string
	// Collect both file and directory names (excluding hidden files).
	for _, name := range entries {
		// name := entry.Name()

		if !strings.HasPrefix(name, ".") { // Exclude hidden files.
			names = append(names, name)
		} else {
			hiddenNames = append(hiddenNames, name)
		}
	}
	if flagsToUse.Flag_a {
		hiddenNames = append(hiddenNames, ".")
		hiddenNames = append(hiddenNames, "..")
		names = append(names, hiddenNames...)
		//AA
	}
	// Sort the names as if they were lowercase but display the actual values.
	sort.SliceStable(names, func(i, j int) bool {
		if flagsToUse.Flag_r {
			// If Flag_r is true, reverse the order.
			return strings.ToLower(names[i]) > strings.ToLower(names[j])
		}
		return strings.ToLower(names[i]) < strings.ToLower(names[j])
	})

	return names, hiddenNames, nil
}

func GetFileType(CWD string, names []string, args []string, flagsToUse Flags, SupCount int) ([]string, []string, []string, error) {
	var fileTypes []string
	var fileColor []string
	for _, name := range names {
		// Join the subfolder path and file name to get the full file path.
		filePath := filepath.Join(CWD, name)

		// Get file information
		// fileInfo, err := os.Lstat(filePath)
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			fmt.Println(len(os.Args))
			fmt.Println("iiiiii", err)

			return nil, nil, args, err
		}
		// List of common graphic image file extensions.
		imageExtensions := ".jpg .jpeg .png .gif .bmp .webp .svg .tiff .ico"
		// List of common archive file extensions.
		archiveExtensions := ".zip .tar .gz .bz2 .xz .rar .7z"

		// Check if it's a directory --Blue
		if fileInfo.IsDir() {
			fileTypes = append(fileTypes, "Directory")
			fileColor = append(fileColor, MyColor("Directory"))

			if flagsToUse.Flag_R {
				args = append(args, filePath)
			}

		} else if fileInfo.Mode()&0111 != 0 {
			// Check if it's an executable file --Green
			fileTypes = append(fileTypes, "Executable")
			fileColor = append(fileColor, MyColor("Executable"))

		} else if fileInfo.Mode()&os.ModeSymlink != 0 {
			// Check if it's a symbolic link --Cyan
			fileTypes = append(fileTypes, "Link")
			fileColor = append(fileColor, MyColor("Link"))

		} else if strings.Contains(imageExtensions, strings.ToLower(FindExt(filePath))) {
			// Check if it's an Image --Pink
			fileTypes = append(fileTypes, "Image")
			fileColor = append(fileColor, MyColor("Image"))

		} else if strings.Contains(archiveExtensions, strings.ToLower(FindExt(filePath))) {
			// Check if it's an archive file --Red
			fileTypes = append(fileTypes, "Archive")
			fileColor = append(fileColor, MyColor("Archive"))

		} else {
			// Default: Unrecognized type
			fileTypes = append(fileTypes, "Unrecognized")
			fileColor = append(fileColor, MyColor("Unrecognized"))

		}
	}

	return fileTypes, fileColor, args, nil
}

func FindExt(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

func MyPath() (string, error) {
	CWD, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return CWD, nil
}

// func DisplayInfo(names []string, color []string) {

// 	// Print the original case names.
// 	for i, name := range names {
// 		if color[i] == "" {
// 			if i == 0 {
// 				fmt.Printf("%v ", name)
// 			} else if i == len(names)-1 {
// 				fmt.Printf(" %v", name)
// 			} else {
// 				fmt.Printf(" %v ", name)
// 			}
// 		} else {
// 			if i == 0 {
// 				fmt.Printf("%v%v%v ", color[i], name, "\033[0m") // first one
// 			} else if i == len(names)-1 {
// 				fmt.Printf(" %v%v%v", color[i], name, "\033[0m") // last one
// 			} else {
// 				fmt.Printf(" %v%v%v ", color[i], name, "\033[0m") // the rest
// 			}
// 		}

// 	}
// 	fmt.Println()

// }
