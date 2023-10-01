package get

type File struct {
	CWD         string
	Names       []string
	HiddenNames []string
	Types       []string
	FileColor   []string
	NotThere    bool
	// IsSubFolder bool
}

type Flags struct {
	Flag_l bool
	Flag_R bool
	Flag_a bool
	Flag_r bool
	Flag_t bool
}
