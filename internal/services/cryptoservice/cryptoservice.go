package cryptoservice

import (
	"errors"

	"github.com/splashk1e/jet/internal/config"
)

type ICryptoService interface {
	Decrypt(text string) (string, error)
	Encrypt(text string) (string, error)
}

func GetCryptoServiceByType(typ string, cfg config.Config) (ICryptoService, error) {
	switch typ {
	case "aes":
		return NewCryptoAesService(cfg.Key), nil
	default:
		return nil, errors.New("wrong service type")
	}
}
