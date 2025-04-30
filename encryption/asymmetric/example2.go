package asymmetric

/*
func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	message := []byte("Important message to sign")

	hasher := sha256.New()
	hasher.Write(message)
	hashed := hasher.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Signature:", base64.StdEncoding.EncodeToString(signature))

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, signature)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
	} else {
		fmt.Println("Signature verified successfully.")
	}
}

// bla bla bla
*/
