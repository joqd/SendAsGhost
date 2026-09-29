package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"SendAsGhost/internal/config"
	"SendAsGhost/internal/handlers"
	"SendAsGhost/internal/handlers/middlewares"
	"SendAsGhost/internal/infrastructure/db"
	"SendAsGhost/internal/infrastructure/telegram"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	conf := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pool, err := db.NewPostgres(ctx, conf.DB.URI)
	if err != nil {
		return err
	}
	defer pool.Close()

	bot, err := telegram.NewBot(conf)
	if err != nil {
		return err
	}

	middlewares.Register(bot, pool)
	handlers.New(pool).Register(bot)

	go func() {
		<-ctx.Done()
		bot.Stop()
	}()

	log.Println("Polling...")
	bot.Start()

	return nil
}