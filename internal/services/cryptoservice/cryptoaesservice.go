package cryptoservice

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"
)

type CryptoAesService struct {
	key []byte
}

func NewCryptoAesService(key []byte) *CryptoAesService {
	return &CryptoAesService{
		key: key,
	}
}
func (s *CryptoAesService) Encrypt(text string) (string, error) {

	cipher, err := aes.NewCipher(s.key)
	if err != nil {
		return "", err
	}

	paddedText := pkcs7Pad([]byte(text), aes.BlockSize)

	out := make([]byte, len(paddedText))

	for i := 0; i < len(paddedText); i += aes.BlockSize {
		cipher.Encrypt(out[i:i+aes.BlockSize], paddedText[i:i+aes.BlockSize])
	}

	return hex.EncodeToString(out), nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize

	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func (s *CryptoAesService) Decrypt(text string) (string, error) {
	ciphertext, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}

	cipher, err := aes.NewCipher(s.key)
	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of block size")
	}

	plain := make([]byte, len(ciphertext))

	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		cipher.Decrypt(plain[i:i+aes.BlockSize], ciphertext[i:i+aes.BlockSize])
	}

	unpadded, err := pkcs7Unpad(plain)
	if err != nil {
		return "", err
	}

	return string(unpadded), nil
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("data length is zero")
	}

	padding := data[length-1]
	padLength := int(padding)

	if padLength > length || padLength > aes.BlockSize {
		return nil, errors.New("invalid padding")
	}

	return data[:length-padLength], nil
}
