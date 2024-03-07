package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgreSQLConfig
	Server   ServerConfig
	Cors     CorsConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type CorsConfig struct {
	AllowOrigins string
}

type PostgreSQLConfig struct {
	User            string
	Password        string
	DatabaseName    string
	Host            string
	SSLMode         string
	Port            string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type JWTConfig struct {
	AccessTokenSecret     string
	AccessTokenExpiresIn  int
	RefreshTokenSecret    string
	RefreshTokenExpiresIn int
}

func SetupConfig() *Config {
	configPath := "config-development"
	v, err := LoadConfig(configPath, "yml")
	if err != nil {
		fmt.Printf("Error in load config %v", err)
	}

	cfg, err := ParseConfig(v)

	if err != nil {
		fmt.Printf("Error in parse config %v", err)
	}

	return cfg

}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath("./config")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}
