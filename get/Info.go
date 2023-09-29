package get

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func GetInfo(cwd string) ([]string, error) {
	// Open the current working directory.
	dir, err := os.Open(cwd)
	if err != nil {
		fmt.Println("Error2:", err)
		return nil, err
	}
	defer dir.Close()

	// Read the contents of the directory.
	entries, err := dir.ReadDir(0) // Read all entries in the directory.
	if err != nil {
		fmt.Println("Error3:", err)
		return nil, err
	}

	// Create a slice to store the names of both files and directories.
	var names []string

	// Collect both file and directory names (excluding hidden files).
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasPrefix(name, ".") { // Exclude hidden files.
			names = append(names, name)
		}
	}

	// Sort the names as if they were lowercase but display the actual values.
	sort.SliceStable(names, func(i, j int) bool {
		return strings.ToLower(names[i]) < strings.ToLower(names[j])
	})
	return names, nil
}

// func GetCWD() (string, error) {
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		fmt.Println("Error5:", err)
// 		return "", err
// 	}
// 	return cwd, nil
// }

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
				fmt.Printf("%v %v %v", color[i], name, "\033[0m")
			} else if i == len(names)-1 {
				fmt.Printf(" %v %v %v ", color[i], name, "\033[0m")
			} else {
				fmt.Printf(" %v %v%v", color[i], name, "\033[0m")
			}
			// Color("reset")
		}

	}
	fmt.Println()

}

func GetFileType(filePaths []string) ([]string, []string, error) {
	var fileTypes []string
	var fileColor []string

	for _, filePath := range filePaths {
		// Get file information
		fileInfo, err := os.Lstat(filePath)
		if err != nil {
			return nil, nil, err
		}
		// List of common graphic image file extensions.
		imageExtensions := ".jpg .jpeg .png .gif .bmp .webp .svg .tiff .ico"
		// List of common archive file extensions.
		archiveExtensions := ".zip .tar .gz .bz2 .xz .rar .7z"

		// Check if it's a directory --Blue
		if fileInfo.IsDir() {
			fileTypes = append(fileTypes, "Directory")
			fileColor = append(fileColor, Color("Directory"))

		} else if fileInfo.Mode()&0111 != 0 {
			// Check if it's an executable file --Green
			fileTypes = append(fileTypes, "Executable")
			fileColor = append(fileColor, Color("Executable"))

		} else if fileInfo.Mode()&os.ModeSymlink != 0 {
			// Check if it's a symbolic link --Cyan
			fileTypes = append(fileTypes, "Link")
			fileColor = append(fileColor, Color("Link"))

		} else if strings.Contains(imageExtensions, strings.ToLower(FindExt(filePath))) {
			// Check if it's an Image --Pink
			fileTypes = append(fileTypes, "Image")
			fileColor = append(fileColor, Color("Image"))

		} else if strings.Contains(archiveExtensions, strings.ToLower(FindExt(filePath))) {
			// Check if it's an archive file --Red
			fileTypes = append(fileTypes, "Archive")
			fileColor = append(fileColor, Color("Archive"))

		} else {
			// Default: Unrecognized type
			fileTypes = append(fileTypes, "Unrecognized")
			fileColor = append(fileColor, "")

		}
	}

	return fileTypes, fileColor, nil
}

func FindExt(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}
