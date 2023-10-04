package get

import (
	"strings"
)

// MyColor returns the ANSI escape sequence for the specified color based on the file type.
func MyColor(fileType string) string {
	var colors = map[string]string{
		"reset": "\033[0m",        // Reset color.
		"red":   "\033[31m",       // Red color.
		"green": "\033[32;1m",     // Green color with bold font.
		"blue":  "\033[34;1m",     // Blue color with bold font.
		"cyan":  "\033[36;1m",     // Cyan color with bold font.
		"pink":  "\033[38;5;206m", // Pink color.

		// You can add more colors here as needed.
	}

	// Define a variable for the color.
	var color string

	// Use a switch statement to set the color based on the file type.
	switch fileType {
	case "Directory":
		color = "Blue" // Set color to Blue for directories.
	case "Executable":
		color = "Green" // Set color to Green for executable files.
	case "Link":
		color = "Cyan" // Set color to Cyan for symbolic links.
	case "Image":
		color = "Pink" // Set color to Pink for image files.
	case "Archive":
		color = "Red" // Set color to Red for archive files.
	default:
		color = "reset" // Set a default color or handle unknown types.
	}

	color = strings.ToLower(color)
	if color1, ok := colors[color]; ok {
		return color1 // Return the ANSI escape sequence for the specified color.
	}

	return "" // Return an empty string for unrecognized colors.
}
