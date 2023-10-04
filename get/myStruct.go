package get

// File represents information about a file or directory.
type File struct {
	CWD              string   // Change Working Directory: The path of the file or directory.
	FileName         string   // The name of the file or folder.
	Names            []string // Names of files inside a folder (if it's a folder).
	HiddenNames      []string // Hidden file names inside a folder (if it's a folder).
	Types            []string // Type of files (e.g., Directory, Executable).
	FileColor        []string // Color representation of file types.
	Permission       []string // File permissions from the system.
	Owner            []string // File owner from the system.
	Group            []string // File group from the system.
	Size             []int64  // File sizes in bytes.
	SizeKB           []int    // File sizes in kilobytes.
	ModificationTime []string // Time of last modification.
	Month            []string // Month of last modification.
	Day              []int    // Day of last modification.
	Hardlinks        []int    // Number of hard links.
	SoftLinks        []string // Soft links for symbolic files.
	TotalSize        int      // Sum of all sizes in KB.
	NotThere         bool     // Indicates if the file or folder does not exist.
	NotFolder        bool     // Indicates if it's not a directory or folder.
	AlingSize        []string
}

// Flags represents command-line flags used in the ls command.
type Flags struct {
	Flag_l bool // List with long format.
	Flag_R bool // List recursively directory tree.
	Flag_a bool // List all files, including hidden ones.
	Flag_r bool // List in reverse order.
	Flag_t bool // Sort by time & date.
}
