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
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	message := []byte("Secret message to encrypt")

	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, message, label)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted message:", base64.StdEncoding.EncodeToString(ciphertext))

	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted message:", string(plaintext))
}
