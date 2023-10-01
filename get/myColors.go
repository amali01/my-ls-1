package get

import (
	"strings"
)

// Color returns the ANSI escape sequence for the specified color
func MyColor(fileType string) string {
	var colors = map[string]string{
		"reset": "\033[0m",
		"red":   "\033[31m",
		"green": "\033[32m",
		"blue":  "\033[34m",
		"cyan":  "\033[36m",
		"pink":  "\033[38;5;206m",

		// "yellow": "\033[33m",
		// "purple": "\033[35m",

		// "white":      "\033[37m",
		// "gray":       "\033[90m",
		// "darkred":    "\033[91m",
		// "orange":     "\033[38;5;208m",
		// "gold":       "\033[38;5;220m",
		// "teal":       "\033[38;5;51m",
		// "lavender":   "\033[38;5;183m",
		// "brown":      "\033[38;5;130m",
		// "lightblue":  "\033[38;5;39m",
		// "magenta":    "\033[38;5;200m",
		// "olive":      "\033[38;5;100m",
		// "salmon":     "\033[38;5;203m",
		// "skyblue":    "\033[38;5;111m",
		// "darkpurple": "\033[38;5;53m",
		// "lime":       "\033[38;5;46m",

		// Add more colors here as needed
	}

	// Define a variable for the color.
	var color string

	// Use a switch statement to set the color based on the file type.
	switch fileType {
	case "Directory":
		color = "Blue"
	case "Executable":
		color = "Green"
	case "Link":
		color = "Cyan"
	case "Image":
		color = "Pink"
	case "Archive":
		color = "Red"
	default:
		color = "reset" // Set a default color or handle unknown types.
	}

	color = strings.ToLower(color)
	if color1, ok := colors[color]; ok {
		return color1
	}

	return "" // Return an empty string for unrecognized colors
}
