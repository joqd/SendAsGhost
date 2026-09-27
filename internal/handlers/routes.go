package handlers

import "gopkg.in/telebot.v4"



func RegisterRoutes(b *telebot.Bot) {
	b.Handle("/ping", pingHandler)
}