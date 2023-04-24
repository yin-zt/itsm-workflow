package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AesTool struct {
	Key       []byte
	BlockSize int
}

func NewAesTool(key []byte, blockSize int) *AesTool {
	return &AesTool{Key: key, BlockSize: blockSize}
}

// Pad plaintext with PKCS#7 padding
func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func (at *AesTool) Encrypt(src []byte) ([]byte, error) {
	var resp any
	// Define key and plaintext
	plaintext := []byte(src)

	// Create AES cipher block
	block, err := aes.NewCipher(at.Key)
	if err != nil {
		resp = err
		panic(resp)
	}

	// Create CBC cipher mode
	iv := at.Key
	mode := cipher.NewCBCEncrypter(block, iv)

	// Pad plaintext to match block size
	paddedPlaintext := pad(plaintext, block.BlockSize())

	// Encrypt padded plaintext
	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	//fmt.Println(encoded)
	return []byte(encoded), nil
}
