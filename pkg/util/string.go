package util

import (
	"math/rand"
	"time"
)

// 生成8位hash盐值
func GenerateSalt() string {
	return RandomString(8)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString returns a random string of the specified length.
func RandomString(length int) string {
	return randomString(length, false)
}

// randomString returns a random string of the specified length.
func randomString(length int, includeNumbers bool) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if includeNumbers {
		letters = append(letters, []byte("0123456789")...)
	}
	return string(randomBytes(length, letters))
}

// randomBytes returns a random byte array of the specified length.
func randomBytes(length int, letters []byte) []byte {
	b := make([]byte, length)
	for i := range b {
		if letters == nil {
			b[i] = byte(randomInt(0, 256))
		} else {
			b[i] = letters[randomInt(0, len(letters))]
		}
	}
	return b
}

// RandomInt returns a random integer in the range [min, max).
func RandomInt(min, max int) int {
	return randomInt(min, max)
}

// randomInt returns a random integer in the range [min, max).
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
