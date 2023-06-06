package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

func DecryptBody(encryptedBody []byte) ([]byte, error) {

	secret := os.Getenv("SECRET")

	key := []byte(secret)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decodedBody, err := base64.URLEncoding.DecodeString(string(encryptedBody))
	if err != nil {
		return nil, err
	}

	decryptedBody := make([]byte, len(decodedBody)-aes.BlockSize)
	iv := decodedBody[:aes.BlockSize]
	encrypted := decodedBody[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decryptedBody, encrypted)

	return decryptedBody, nil
}
