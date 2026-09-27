package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Bot Bot
		Proxy string
	}

	Bot struct {
		Token string
	}
)

var (
	configInstance *Config
	once sync.Once
)

func Load() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		configInstance = &Config{
			Bot: Bot{
				Token: os.Getenv("BOT_TOKEN"),
			},
			Proxy: os.Getenv("PROXY"),
		}
	})

	return configInstance
}