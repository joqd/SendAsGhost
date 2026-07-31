package main

import (
	"SendAsGhost/internal/config"
	"SendAsGhost/internal/infrastructure/poller"
)

func main() {
	conf := config.Load()
	poller.NewPoller(conf).Start()
}