package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Slack struct {
	params params
	url    string
}

type params struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

// NewSlack はSlackにIncoming Webhooksで通知します
func NewSlack(url string, text string, username string, iconEmoji string, iconURL string, channel string) *Slack {
	p := params{
		Text:      text,
		Username:  username,
		IconEmoji: iconEmoji,
		IconURL:   iconURL,
		Channel:   channel,
	}

	return &Slack{
		params: p,
		url:    url,
	}
}

func (s *Slack) Send() {
	params, _ := json.Marshal(s.params)
	resp, err := http.PostForm(
		s.url,
		url.Values{"payload": {string(params)}},
	)
	defer resp.Body.Close()

	if err != nil {
		log.Print(err)
		return
	}
}
