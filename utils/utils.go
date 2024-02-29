package utils

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveKey(passphrase string, salt []byte) []byte {
	return pbkdf2.Key([]byte(passphrase), salt, 4096, 32, sha256.New)
}
