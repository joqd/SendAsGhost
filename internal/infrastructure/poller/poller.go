package poller

import (
	"SendAsGhost/internal/config"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	tele "gopkg.in/telebot.v4"
)


type Poller struct {
	bot *tele.Bot
}

func NewPoller(conf *config.Config) *Poller {
	pref := tele.Settings{
		Token:  conf.Bot.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	
	if conf.Proxy != "" {
		log.Println("Parsing proxy")

		proxyURL, err := url.Parse("http://127.0.0.1:10808")
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
			},
		}

		log.Println("Setting proxy")
		pref.Client = client
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	return &Poller{
		bot: b,
	}
}

func (p *Poller) Start() {
	log.Println("Polling...")
	p.bot.Start()
}