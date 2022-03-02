package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	PORT        string `mapstructure:"PORT"`
	Environment string `mapstructure:"APP_ENV"`
	MongoDbName string `mapstructure:"MONGO_DB_NAME"`
	MongoDbUrl  string `mapstructure:"MONGO_DB_URL"`
	Salt        int    `mapstructure:"SALT"`
	JwtSecret   string `mapstructure:"JWT_SECRET"`
	JwtExpiry   int    `mapstructure:"JWT_EXPIRY"`
}

func Get(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var configuration Config
	err = viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return &configuration, nil
}