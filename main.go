package main

import (
	"github.com/brkss/volt/cmd"
)

func main() {

	cmd.Execute()

	/*
		// Reading plaintext from a file
		plaintext, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		// Encrypting the plaintext
		ciphertext, err := crypting.Encrypt(plaintext, passphrase)
		if err != nil {
			log.Fatal("encrypt : ", err)
		}

		// // Writing the encrypted data to a file
		// if err := ioutil.WriteFile("encrypted.txt", ciphertext, 0644); err != nil {
		// 	log.Fatal(err)
		// }

		// Decrypting the ciphertext
		decrypted, err := crypting.Decrypt(ciphertext, passphrase)
		if err != nil {
			log.Fatal(err)
		}

		// // Writing the decrypted data to a file
		// if err := ioutil.WriteFile("decrypted.txt", decrypted, 0644); err != nil {
		// 	log.Fatal(err)
		// }
	*/
}
