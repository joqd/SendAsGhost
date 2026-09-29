package main

import (
	"SendAsGhost/internal/config"
	"SendAsGhost/internal/infrastructure/db"
	"SendAsGhost/internal/infrastructure/poller"
	"context"
	"log"
)

func main() {
	conf := config.Load()

	ctx := context.Background()
	pool, err := db.NewPostgres(ctx, conf.DB.URI)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	poller.NewPoller(conf).Start()
}