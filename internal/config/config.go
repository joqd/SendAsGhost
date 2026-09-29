package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		DB DB
		Bot Bot
		Proxy string
	}

	DB struct {
		URI string
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
			DB: DB{
				URI: os.Getenv("PG_URI"),
			},
			Bot: Bot{
				Token: os.Getenv("BOT_TOKEN"),
			},
			Proxy: os.Getenv("PROXY"),
		}
	})

	return configInstance
}