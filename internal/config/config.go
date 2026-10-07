package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type (
    Config struct {
        Bot Bot
        Proxy string
		Postgres Postgres
    }

    Postgres struct {
        DB string
		User string
		Password string
		Host string
		Port int
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

		port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		if err != nil {
			log.Fatal("Postgres port must be int")
		}

        configInstance = &Config{
            Postgres: Postgres{
                DB: os.Getenv("POSTGRES_DB"),
				User: os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				Host: os.Getenv("POSTGRES_HOST"),
				Port: port,
            },
            Bot: Bot{
                Token: os.Getenv("BOT_TOKEN"),
            },
            Proxy: os.Getenv("PROXY"),
        }
    })

    return configInstance
}