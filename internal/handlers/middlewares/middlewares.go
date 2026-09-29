package middlewares

import (
	"github.com/jackc/pgx/v5/pgxpool"
	tele "gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

func Register(bot *tele.Bot, pool *pgxpool.Pool) {
	bot.Use(middleware.Logger())
	bot.Use(middleware.AutoRespond())
	bot.Use(EnsureUser(pool))
}