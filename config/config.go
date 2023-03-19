package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Addr            string
	WriteTimeout    int
	ReadTimeout     int
	GraceFulTimeout int
}

type DBConfig struct {
	Name            string
	Host            string
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifetime int
}

func InitConfig() Config {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	var configuration Config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return configuration
}
