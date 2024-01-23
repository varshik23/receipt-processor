package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(receipt receipt) string {
	data := []byte(fmt.Sprintf("%+v", receipt))

	// Calculate the SHA-256 hash
	hash := sha256.Sum256(data)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash[:])

	return hashString
}
