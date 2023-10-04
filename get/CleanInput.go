package get

import (
	"os"
	"strings"
)

// CleanedInputargs removes empty strings from a slice of input arguments.
func CleanedInputargs(inputargs []string) []string {
	cleanedInputargs := make([]string, 0, len(inputargs))
	for _, arg := range inputargs {
		if arg != "" {
			cleanedInputargs = append(cleanedInputargs, arg)
		}
	}
	return cleanedInputargs
}

// TrimFliePath trims the original file path to remove the main root prefix.
func TrimFliePath(originalFliePath string, mainRoot string) string {
	trimmedFliePath := strings.TrimPrefix(originalFliePath, mainRoot)
	return trimmedFliePath
}

// FindExt finds the file extension in a given path.
func FindExt(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

// MyPath returns the current working directory.
func MyPath() (string, error) {
	CWD, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return CWD, nil
}

// GetUpperPath gets the parent folder path of the specified directory.
func GetUpperPath(path string) string {
	pathInRune := []rune(path)
	var dashCounter int = 0

	for k := 0; k < len(pathInRune); k++ {
		if pathInRune[k] == '/' {
			dashCounter++
		}
	}

	if dashCounter == 1 {
		return "/"
	}

	for k := 0; k < len(path); k++ {
		if path[len(path)-1:] == "/" {
			path = path[:len(path)-1]
			break
		} else {
			path = path[:len(path)-1]
		}
	}

	return path
}

// IfContains checks if a slice of strings contains a specific string.
func IfContains(RootNames []string, arg string) bool {
	for _, name := range RootNames {
		if name == arg {
			return true
		}
	}
	return false
}

// // Checks if folder or file is hidden
// func IsHidden(filename string) bool {
// 	if filename[0:1] == "." {
// 		return true
// 	}
// 	return false
// }
