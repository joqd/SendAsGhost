package ytdlp

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	urlRe = regexp.MustCompile(`https?://[^\s]+`)
	idRe  = regexp.MustCompile(`^[A-Za-z0-9_-]{11}$`)
)

// ExtractVideoURL finds the first YouTube video link in text and returns
// a canonical watch URL. Playlist parameters are dropped.
func ExtractVideoURL(text string) (string, bool) {
	for _, raw := range urlRe.FindAllString(text, -1) {
		u, err := url.Parse(raw)
		if err != nil {
			continue
		}

		host := strings.TrimPrefix(strings.ToLower(u.Hostname()), "www.")
		segments := strings.Split(strings.Trim(u.Path, "/"), "/")

		var id string
		switch host {
		case "youtu.be":
			id = segments[0]
		case "youtube.com", "m.youtube.com", "music.youtube.com":
			if u.Path == "/watch" {
				id = u.Query().Get("v")
			} else if len(segments) == 2 &&
				(segments[0] == "shorts" || segments[0] == "live" || segments[0] == "embed") {
				id = segments[1]
			}
		}

		if idRe.MatchString(id) {
			return "https://www.youtube.com/watch?v=" + id, true
		}
	}
	return "", false
}