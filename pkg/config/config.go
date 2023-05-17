package config

import (
	"log"

	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	RPCURL        string `mapstructure:"ETH_RPC_URL"`
	PRIVATEKEY    string `mapstructure:"PRIVATE_KEY"`
	GASLIMIT      uint64 `mapstructure:"GAS_LIMIT"`
	INCREASEGAS   int64  `mapstructure:"INCREASE_GAS"`
	TOKEN         string `mapstructure:"TOKEN_FILE_NAME"`
	TOKENADDR     string `mapstructure:"TOKEN_CONTRACT_ADDR"`
	TOKENDECIMALS int    `mapstructure:"TOKEN_CONTRACT_DECIMALS"`
	AIRDROP       string `mapstructure:"AIRDROP_FILE_NAME"`
	AIRDROPADDR   string `mapstructure:"AIRDROP_CONTRACT_ADDR"`
	AIRDROPFILE   string `mapstructure:"AIRDROP_FILES"`
	AIRDROPPER    int    `mapstructure:"AIRDROP_PER"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	processPath, err := ProcessPath()
	if err != nil {
		log.Fatal("Error get process path", err)
		return
	}

	// setting viper path and file name.
	viper.AddConfigPath(processPath)
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading env file", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ProcessPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return filepath.Dir(ex), nil
}
