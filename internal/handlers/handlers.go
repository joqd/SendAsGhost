package handlers

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	tele "gopkg.in/telebot.v4"
)

type Handler struct {
	ctx   context.Context
	pool  *pgxpool.Pool
}

func New(ctx context.Context, pool *pgxpool.Pool) *Handler {
	return &Handler{ctx: ctx, pool: pool}
}

func (h *Handler) Register(bot *tele.Bot) {
	bot.Handle("/ping", h.Ping)
}