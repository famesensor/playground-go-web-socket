package config

import (
	"log"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Env      string
	Postgres Postgres
}

type App struct {
	Name    string
	Version string
	Port    string
}

type Postgres struct {
	Username    string
	Password    string
	Host        string
	Port        string
	BDName      string
	Schema      string
	SslMode     string
	TimeZone    string
	MaxIdle     int
	MaxOpen     int
	MaxLifeTime time.Duration
	MaxIdleTime time.Duration
}

func NewConfig() *Config {
	viper.SetConfigName("config") // config.yaml, config.json, etc.
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // Look for config in current directory
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // Read environment variables

	if err := viper.ReadInConfig(); err != nil {
		log.Println("[CONFIG] config file not found, using defaults")
		return nil
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("[CONFIG] config file changed:", e.Name)
	})
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Println("[CONFIG] config Error:", err)
		return nil
	}

	return &config
}
