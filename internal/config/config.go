package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host                string `yaml:"host"`
	Port                string `yaml:"port"`
	SoonExpiredWarning  time.Duration
	SoonExpiredCritical time.Duration
	FilePath            string `yaml:"file_path"`
	FilePathDev         string `yaml:"file_path_dev"`
	Key                 []byte `yaml:"aes256_key"`
}

func (c *Config) InitConfig() error {
	return cleanenv.ReadConfig("config.yaml", c)
}
