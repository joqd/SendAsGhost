package main

import (
	"context"
	"fmt"
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

	// Connecting to Postgres
	pg := conf.Postgres
	uri := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		pg.User, pg.Password, pg.Host, pg.Port, pg.DB,
	)
	pool, err := db.NewPostgres(ctx, uri)
	if err != nil {
		return err
	}
	defer pool.Close()

	// Setup handlers/middlewares
	bot, err := telegram.NewBot(conf)
	if err != nil {
		return err
	}

	middlewares.Register(bot, pool)
	handlers.New(ctx, pool).Register(bot)


	// Polling
	go func() {
		<-ctx.Done()
		bot.Stop()
	}()

	log.Println("Polling...")
	bot.Start()

	return nil
}