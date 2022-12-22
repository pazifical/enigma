package internal

type File struct {
	Data []byte
	Path string
}

type EncryptedFile File

type UnencryptedFile File

type DecryptedFile File
