package handlers

import (
	tele "gopkg.in/telebot.v4"
)

func pingHandler(c tele.Context) error {
	return c.Send("Pong")
}