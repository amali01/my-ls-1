package Structs

type file struct {
	CWD       string
	Names     []string
	Types     []string
	fileColor []string
	NotThere  bool
}

type Flags struct {
	Flag_l bool
	Flag_R bool
	Flag_a bool
	Flag_r bool
	Flag_t bool
}
