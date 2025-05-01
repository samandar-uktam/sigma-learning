package signature

/*
func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	message := []byte("this is a secret message")
	hash := sha256.Sum256(message)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ ECDSA Signature created.")

	isValid := ecdsa.Verify(publicKey, hash[:], r, s)
	if isValid {
		fmt.Println("✅ ECDSA Signature verified.")
	} else {
		fmt.Println("❌ ECDSA Signature verification failed.")
	}
}
*/
