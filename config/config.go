// Package config  will be basis of environment configs in future
package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

//Config is struct for config
type Config struct {
	Port  string
	Debug bool
	DB    DBConfig
}

//DBConfig is part is package config
type DBConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Log      bool
}

var config *Config

//New returns a config if exists otherwise returns new
func New() *Config {
	if config == nil {
		config = &Config{
			Port: getEnv("PORT", "9095"),

			Debug: getEnvBool("DEBUG", true),
			DB: DBConfig{
				Username: getEnv("DB_USERNAME", ""),
				Password: getEnv("DB_PASSWORD", ""),
				Name:     getEnv("DB_NAME", "generic_service_local"),
				Host:     getEnv("DB_HOST", "localhost"),
				Log:      getEnvBool("DB_LOG", getEnvBool("DEBUG", true)),
			},
		}
	}
	return config
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		bool, err := strconv.ParseBool(value)
		if err != nil {
			log.Error(err)
		}
		return bool
	}

	return defaultVal
}

// SetEnvFromDotEnv sets the environment defined in dotenv
func SetEnvFromDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. Rename '.env.sample' to '.env' and add your environment data  : %w", err)
	}
	log.Info("Using .env environment")
}
