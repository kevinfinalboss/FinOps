package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const webhookURL = "https://discord.com/api/webhooks/1157394403594346556/MypTfWBNLwQ4Cwgn0YtPza_HwBw9WvfFCV1T6IXPJwgALwOJqqJu2Rr2Z_gvZwxN7ATg"

type DiscordEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
	Color       int    `json:"color"`
	Footer      struct {
		Text string `json:"text"`
	} `json:"footer"`
	Image struct {
		URL string `json:"url"`
	} `json:"image"`
	Fields []struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Inline bool   `json:"inline"`
	} `json:"fields"`
}

func SendLoginWebhook(username, email, ip string) error {
	embed := DiscordEmbed{
		Title:       "Novo Login",
		Description: "Um usuário logou no sistema.",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0x00ff00,
		Footer: struct {
			Text string `json:"text"`
		}{
			Text: "FinOps Login System",
		},
		Image: struct {
			URL string `json:"url"`
		}{
			URL: "https://github.com/kevinfinalboss/StoreOps/blob/master/screenshots/Logo.jpg?raw=true",
		},
		Fields: []struct {
			Name   string `json:"name"`
			Value  string `json:"value"`
			Inline bool   `json:"inline"`
		}{
			{
				Name:   "Nome",
				Value:  username,
				Inline: true,
			},
			{
				Name:   "Email",
				Value:  email,
				Inline: true,
			},
			{
				Name:   "Data/Horário",
				Value:  time.Now().Format("02/01/2006 15:04:05"),
				Inline: true,
			},
			{
				Name:   "IP",
				Value:  ip,
				Inline: true,
			},
		},
	}

	body := map[string]interface{}{
		"embeds": []DiscordEmbed{embed},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = http.Post(webhookURL, "application/json", bytes.NewBuffer(bodyBytes))
	return err
}
