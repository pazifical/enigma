package encryption

type Writer struct {
	filepath string
}

func NewWriter(filepath string) Writer {
	return Writer{filepath: filepath}
}
