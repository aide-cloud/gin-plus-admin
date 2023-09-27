package util

// EncryptPassword ...
func EncryptPassword(password, salt string) (string, error) {
	aes, err := NewAesCipher()
	if err != nil {
		return "", err
	}
	return aes.EncryptAesBase64([]byte(password + salt))
}
