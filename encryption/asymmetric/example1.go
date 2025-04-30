package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	// Generate a new RSA key pair (2048-bit).
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	// The message to encrypt.
	message := []byte("Secret message to encrypt")

	// Encrypt the message with the public key using RSA-OAEP.
	// OAEP provides better security than raw PKCS#1 v1.5 padding.
	label := []byte("") // OAEP label can be empty or used for additional context
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, message, label)
	if err != nil {
		log.Fatal(err)
	}
	// Print ciphertext in Base64 for readability.
	fmt.Println("Encrypted message:", base64.StdEncoding.EncodeToString(ciphertext))

	// Decrypt the ciphertext with the private key.
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted message:", string(plaintext))
}
