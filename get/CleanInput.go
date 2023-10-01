package get

import (
	"strings"
)

func CleanedInputargs(inputargs []string) []string {
	cleanedInputargs := make([]string, 0, len(inputargs))
	for _, arg := range inputargs {
		if arg != "" {
			cleanedInputargs = append(cleanedInputargs, arg)
		}
	}
	return cleanedInputargs
}

func TrimPath(originalPath string, IsSubFolder bool) string {
	// Split the path into components using the forward slash as a separator
	components := strings.Split(originalPath, "/")
	trimmedPath := ""

	if IsSubFolder {
		// Extract the last two component
		trimmedPath = strings.Join(components[len(components)-2:], "/")

	} else {
		// Extract the last component
		trimmedPath = components[len(components)-1]

	}

	return trimmedPath
}
