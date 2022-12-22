package encryption

import "github.com/TwoWaySix/enigma/internal"

type Processor struct {
	encryptionKey string
}

func NewProcessor(encryptionKey string) Processor {
	return Processor{encryptionKey: encryptionKey}
}

func (p *Processor) Encrypt(data internal.UnencryptedFile) internal.EncryptedFile {
	// TODO: implement
	return internal.EncryptedFile(data)
}
