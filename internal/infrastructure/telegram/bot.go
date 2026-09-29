package telegram

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	tele "gopkg.in/telebot.v4"

	"SendAsGhost/internal/config"
)

func NewBot(conf *config.Config) (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  conf.Bot.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	if conf.Proxy != "" {
		proxyURL, err := url.Parse(conf.Proxy)
		if err != nil {
			return nil, fmt.Errorf("parse proxy url: %w", err)
		}

		pref.Client = &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
			},
		}
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("create bot: %w", err)
	}

	return bot, nil
}