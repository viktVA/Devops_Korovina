package cfg

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"os"
)

const (
	keyPath = "./internal/cfg/ec_key.pem"
	cfgPath = "./internal/cfg/config.json"
)

type Config struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func NewCfg() (*Config, error) {
	jsonFile, err := os.Open(cfgPath)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.NewDecoder(jsonFile).Decode(&config)

	pem, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	config.PrivateKey, err = jwt.ParseECPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, err
	}
	config.PublicKey = &config.PrivateKey.PublicKey

	return &config, err
}
