package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Strategy      string `mapstructure:"STRATEGY"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	LimitIp       int    `mapstructure:"LIMIT_IP"`
	LimitToken    int    `mapstructure:"LIMIT_TOKEN"`
	BlockTime     int    `mapstructure:"BLOCK_TIME"`
}

func LoadConfig() (*Config, error) {
	var cfg *Config
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
