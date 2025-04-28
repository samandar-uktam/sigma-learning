package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// encrypt encrypts the given plaintext using AES-GCM
func encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// GCM mode (Galois Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Create a random nonce (number used once)
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the data
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// decrypt decrypts the given ciphertext using AES-GCM
func decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// Split nonce and actual encrypted data
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {
	// Secret key must be either 16, 24, or 32 bytes (AES-128, AES-192, AES-256)
	key := []byte("thisis32bitlongpassphraseimusing!") // 32 bytes => AES-256

	message := "Hello, this is a secret message!"

	// Encrypt the message
	encrypted, err := encrypt([]byte(message), key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %x\n", encrypted)

	// Decrypt the message
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
