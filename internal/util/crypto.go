package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Encrypt(plaintext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		Log().Error(err.Error())
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		Log().Error(err.Error())
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		Log().Error(err.Error())
		return "", err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return fmt.Sprintf("%x", string(ciphertext)), nil // Return as hex
}

func Decrypt(ciphertext, key string) (string, error) {
	bCiphertext, err := hex.DecodeString(ciphertext)
	if err != nil {
		Log().Error(err.Error())
		return "", err
	}
	ciphertext = string(bCiphertext)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		Log().Error(err.Error())
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		Log().Error(err.Error())
		return "", err
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := aesgcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesgcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		Log().Error(err.Error())
		return "", err
	}

	return string(plaintext), nil
}
