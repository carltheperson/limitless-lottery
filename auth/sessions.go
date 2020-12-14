package auth

import (
	"encoding/base64"
	"math/rand"
	"time"
)

const tokenLength = 30

func GenerateSessionToken() string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := generateRandomBytes(tokenLength)
	return base64.URLEncoding.EncodeToString(b)
}

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
