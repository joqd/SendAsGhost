package handlers

import (
	"context"
	"time"

	tele "gopkg.in/telebot.v4"
)

func (h *Handler) Ping(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := h.pool.Ping(ctx); err != nil {
		return c.Send("db is down")
	}
	return c.Send("pong")
}