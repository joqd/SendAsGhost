package main

import (
	"SendAsGhost/internal/config"
	"log"
)

func main() {
	conf := config.Load()
	log.Println(conf.Bot.Token)
}