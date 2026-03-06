package config

import (
	"price-tracker-service/src/domain"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func LoadDatabaseConfig() (*domain.DatabaseConfig, error) {
	var databaseConfig domain.DatabaseConfig
	if err := viper.Unmarshal(&databaseConfig); err != nil {
		return nil, err
	}

	return &databaseConfig, nil
}
