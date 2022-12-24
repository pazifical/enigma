package internal

type Config struct {
	Mode      string
	InputPath string
	Key       string
	OutPath   string
	Paths     []string // TODO: Remove when deprecated
}
