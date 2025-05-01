package signature

/*
func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	message := []byte("this is a secret message")
	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ RSA Signature created.")

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("❌ RSA Signature verification failed:", err)
	} else {
		fmt.Println("✅ RSA Signature verified.")
	}
}
*/
