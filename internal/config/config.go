package config

import (
	"log"

	"github.com/1tsandre/mini-go-backend/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Port int `mapstructure:"port"`
    } `mapstructure:"server"`

    Database struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        User     string `mapstructure:"user"`
        Password string `mapstructure:"password"`
        Name     string `mapstructure:"name"`
        SSLMode  string `mapstructure:"sslmode"`
    } `mapstructure:"database"`

    Redis struct {
        Address  string `mapstructure:"address"`
        Password string `mapstructure:"password"`
        DB       int    `mapstructure:"db"`
    } `mapstructure:"redis"`
}

func LoadConfig() *Config {
    v := viper.New()
    v.SetConfigName("config")
    v.SetConfigType("yaml")
	v.AddConfigPath(".")

    if err := v.ReadInConfig(); err != nil {
        logger.Fatalf("Error reading config: %v", err)
    }

    var cfg Config
    if err := v.Unmarshal(&cfg); err != nil {
        log.Fatalf("Error parsing config: %v", err)
    }

    logger.Infof("Config loaded")
    return &cfg
}
