package file_service

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	str := make([]byte, length)
	for i := range str {
		str[i] = charset[rng.Intn(len(charset))]
	}
	return string(str)
}
