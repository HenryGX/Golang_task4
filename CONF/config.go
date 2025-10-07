package CONF

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

// Config holds the application configuration.
type Config struct {
	JWTSecret string `mapstructure:"jwt_secret"`
	DSN       string `mapstructure:"dsn"`
}

// AppConfig is the application's configuration.
var AppConfig Config
var once sync.Once

// LoadConfig loads configuration from a file.
func LoadConfig() {
	once.Do(func() {
		viper.AddConfigPath("./conf")
		viper.SetConfigName("config")
		viper.SetConfigType("json")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		err := viper.Unmarshal(&AppConfig)
		if err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
	})
}

// GetConfig returns the application configuration.
// This function ensures that the configuration is loaded before returning.
func GetConfig() Config {
	LoadConfig()
	return AppConfig
}
