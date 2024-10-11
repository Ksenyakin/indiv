// pkg/config/config.go
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel                  string
	GRPCPort                  string
	RESTPort                  string
	Database                  DatabaseConfig
	PaymentServiceURL         string
	PaymentServiceAPIKey      string
	NotificationServiceURL    string
	NotificationServiceAPIKey string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
