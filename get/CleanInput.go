package get

import (
	"fmt"
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

// SizeFormat formats the Size field in a slice of File structs with leading spaces.
func SizeFormat(files []File) []File {
	for i := range files {
		// Skip files that are not there
		if files[i].NotThere {
			continue
		}

		// Find the maximum width of numbers in the Size slice
		maxWidth := FindMaxWidth(files[i].Size)

		// Create a new slice to store updated sizes
		updatedSize := make([]string, len(files[i].Size))

		// Update each size with leading spaces and store in updatedSize
		for j, num := range files[i].Size {
			updatedSize[j] = fmt.Sprintf("%*d", maxWidth, num)
		}

		// Assign the updatedSize to the AlingSize field in the File struct
		files[i].AlingSize = updatedSize
	}

	return files
}

// findMaxWidth finds the maximum width of numbers in a slice
func FindMaxWidth(numbers []int64) int {
	maxWidth := 0
	for _, num := range numbers {
		// Calculate the width of the number
		width := NumWidth(num)

		// Update maxWidth if width is greater
		if width > maxWidth {
			maxWidth = width
		}
	}
	return maxWidth
}

// numWidth calculates the width of an int64 number
func NumWidth(num int64) int {
	width := 0
	for num != 0 {
		num /= 10
		width++
	}
	return width
}
