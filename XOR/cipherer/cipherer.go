package cipherer

import (
	"encoding/base64"
	"errors"
)

func Cipher(rawString, secret string) (string, error) {
	if len(secret) == 0 {
		return "", errors.New("ключ не указан")
	}
	encryptedBytes, err := process([]byte(rawString), []byte(secret))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}
func Decipher(cipheredText, secret string) (string, error) {
	if len(secret) == 0 {
		return "", errors.New("ключ не указан")
	}
	cipheredBytes, err := base64.StdEncoding.DecodeString(cipheredText)

	if err != nil {
		return "", errors.New("какая то ошибка, может связано с тем, что неправильная декодировка на бейз 64")
	}

	decryptedBytes, err := process(cipheredBytes, []byte(secret))
	if err != nil {
		return "", err
	}

	return string(decryptedBytes), nil
}
func process(input, secret []byte) ([]byte, error) {
	for i, b := range input {
		input[i] = b ^ secret[i%len(secret)]
	}
	return input, nil
}
