package util

import (
	"math/rand"
	"time"
)

// GenerateSalt 生成8位hash盐值
func GenerateSalt() string {
	return RandomString(8)
}

// RandomString 返回指定长度的随机字符串
func RandomString(length int) string {
	return randomString(length, false)
}

// randomString 返回指定长度的随机字符串
func randomString(length int, includeNumbers bool) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if includeNumbers {
		letters = append(letters, []byte("0123456789")...)
	}
	return string(randomBytes(length, letters))
}

// randomBytes 返回指定长度的随机字节切片
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

// RandomInt 返回指定范围内的随机整数
func RandomInt(min, max int) int {
	return randomInt(min, max)
}

// randomInt 返回指定范围内的随机整数
func randomInt(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return min + r.Intn(max-min)
}
