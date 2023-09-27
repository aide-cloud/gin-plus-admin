package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const (
	// defaultAesKey 默认的aes key 32位
	defaultAesKey = "08d19803fb6e9b1e9ef0830d4a2e148e"
	// defaultAesIv 默认的aes iv 16位
	defaultAesIv = "830d4a2e121da2e1"
)

type (
	AesCipher struct {
		key   []byte
		iv    []byte
		block cipher.Block
	}

	AesCipherInterface interface {
		EncryptAesBase64(in []byte) (string, error)
		DecryptAesBase64(b string) ([]byte, error)
	}

	AesCipherOption func(*AesCipher)
)

// NewAesCipher 创建一个新的AesCipher
func NewAesCipher(opts ...AesCipherOption) (*AesCipher, error) {
	aesExcept := AesCipher{
		key: []byte(defaultAesKey),
		iv:  []byte(defaultAesIv),
	}

	for _, opt := range opts {
		opt(&aesExcept)
	}

	var err error

	aesExcept.block, err = aes.NewCipher(aesExcept.key)
	if err != nil {
		return nil, err
	}
	return &aesExcept, nil
}

func WithAesKey(key [32]byte) AesCipherOption {
	return func(a *AesCipher) {
		a.key = []byte(key[:])
	}
}

func WithAesIv(iv [16]byte) AesCipherOption {
	return func(a *AesCipher) {
		a.iv = []byte(iv[:])
	}
}

// EncryptAesBase64 加密
func (a *AesCipher) EncryptAesBase64(in []byte) (string, error) {
	origData := in
	origData = pkCS5Padding(origData, a.block.BlockSize())
	crypt := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypt也可以
	bm := cipher.NewCBCEncrypter(a.block, a.iv)
	bm.CryptBlocks(crypt, origData)
	var b = base64.StdEncoding.EncodeToString(crypt)
	return b, nil
}

// DecryptAesBase64 解密
func (a *AesCipher) DecryptAesBase64(b string) ([]byte, error) {
	crypt, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return []byte{}, err
	}
	origData := make([]byte, len(crypt))
	bm := cipher.NewCBCDecrypter(a.block, a.iv)
	bm.CryptBlocks(origData, crypt)
	origData = pkCS5UnPadding(origData)
	return origData, nil
}

func pkCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pkCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unPadding 次
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
