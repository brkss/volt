package crypting

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/brkss/volt/utils"
)

func Decrypt(data []byte, passphrase string) ([]byte, error) {
	salt := data[:8]
	ciphertext := data[8:]

	key := utils.DeriveKey(passphrase, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
