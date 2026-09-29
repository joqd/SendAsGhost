package handlers

import tele "gopkg.in/telebot.v4"

func (h *Handler) Register(bot *tele.Bot) {
	bot.Handle("/ping", h.Ping)
}