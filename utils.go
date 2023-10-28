package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func ToString(val any) string {
	return fmt.Sprintf("%v", val)
}

// it creates unique short codes based on urls provided using hashing techniques
func createShortCode(url string) (shortCode string) {
	sha256 := sha256.New()
	sha256.Write([]byte(url))
	checksum := sha256.Sum(nil)
	hashedURL := base64.URLEncoding.EncodeToString(checksum)

	shortCode = string(hashedURL)

	return shortCode
}
