package pkg

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func Hashing(content io.Reader) ([]byte, string) {
	//hash := sha1.New()
	fileContent, _ := io.ReadAll(content)
	hash := []byte(string(fileContent))
	checksum := sha1.Sum(hash)
	hashed := hex.EncodeToString(checksum[:])
	return fileContent, hashed
}
