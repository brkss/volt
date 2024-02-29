package crypting

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/brkss/volt/utils"
)

func Encrypt(data []byte, passphrase string) ([]byte, error) {
	salt := make([]byte, 8) // generate random salt
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}

	key := utils.DeriveKey(passphrase, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return append(salt, ciphertext...), nil
}
