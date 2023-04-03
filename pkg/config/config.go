package config

import (
	"log"

	"github.com/spark8899/airdropErc20/pkg/utils"
	"github.com/spf13/viper"
)

type Config struct {
	RPCURL     string `mapstructure:"ETH_RPC_URL"`
	PRIVATEKEY string `mapstructure:"PRIVATE_KEY"`
	GASLIMIT   int64  `mapstructure:"GAS_LIMIT"`
	TOKEN      string `mapstructure:"TOKEN_FILE_NAME"`
	AIRDROP    string `mapstructure:"AIRDROP_FILE_NAME"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	processPath, err := utils.ProcessPath()
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
