package cryptoservice

import (
	"testing"
)

func TestCryptoAesService_Encrypt(t *testing.T) {
	tests := []struct {
		name    string
		s       *CryptoAesService
		text    string
		want    string
		wantErr bool
	}{
		{
			"Обычный пример",
			NewCryptoAesService([]byte{123, 12, 53, 63, 12, 65, 158, 132, 124, 25, 72, 183, 184, 103, 15, 1, 245, 223, 76, 45, 83, 37, 53, 27, 47, 38, 28, 90, 16, 26, 31, 65}),
			"тестовый пример",
			"b646143efea3baef1cb2a9ddea5a40bd6932a850f9eee0f7fc3b63f2b54c26f3",
			false,
		},
		{
			"Неверный ключ",
			NewCryptoAesService([]byte{123, 12, 53, 63, 12, 65, 158, 132, 124, 25, 72, 183, 184, 103, 15, 1, 245, 223, 76, 45, 83, 37, 53, 27, 47, 38, 28, 90, 16, 26, 31}),
			"тестовый пример",
			"b646143efea3baef1cb2a9ddea5a40bd6932a850f9eee0f7fc3b63f2b54c26f3",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Encrypt(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoAesService.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CryptoAesService.Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCryptoAesService_Decrypt(t *testing.T) {
	tests := []struct {
		name    string
		s       *CryptoAesService
		text    string
		want    string
		wantErr bool
	}{
		{
			"Обычный пример",
			NewCryptoAesService([]byte{123, 12, 53, 63, 12, 65, 158, 132, 124, 25, 72, 183, 184, 103, 15, 1, 245, 223, 76, 45, 83, 37, 53, 27, 47, 38, 28, 90, 16, 26, 31, 65}),
			"b646143efea3baef1cb2a9ddea5a40bd6932a850f9eee0f7fc3b63f2b54c26f3",
			"тестовый пример",
			false,
		},
		{
			"Неверный ключ",
			NewCryptoAesService([]byte{123, 12, 53, 63, 12, 65, 158, 132, 124, 25, 72, 183, 184, 103, 15, 1, 245, 223, 76, 45, 83, 37, 53, 27, 47, 38, 28, 90, 16, 26, 31}),
			"тестовый пример",
			"b646143efea3baef1cb2a9ddea5a40bd6932a850f9eee0f7fc3b63f2b54c26f3",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Decrypt(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoAesService.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CryptoAesService.Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
