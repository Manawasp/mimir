package config

import (
	"bytes"
	"encoding/json"
	"feedback/api/models"
	"net/http"
	"regexp"
)

type WebhookChannel struct {
	Name    string `toml:"channel"`
	Pattern string `toml:"pattern"`
}

type Webhook struct {
	URL      string           `toml:"url"`
	Icon     string           `toml:"icon"`
	Username string           `toml:"username"`
	Channels []WebhookChannel `toml:"channel"`
}

type NotifyMessage struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
	Icon     string `json:"icon"`
}

func (wh *Webhook) Notify(feedback *models.Feedback) {
	// Init notif
	notif := NotifyMessage{
		Text:     createText(feedback),
		Icon:     wh.Icon,
		Username: wh.Username,
	}

	// Set channel if match else default
	for _, h := range wh.Channels {
		if matched, _ := regexp.MatchString(h.Pattern, feedback.URL); matched {
			notif.Channel = h.Name
			break
		}
	}

	// Send notif
	data, _ := json.Marshal(map[string]interface{}{"payload": notif})
	http.Post(wh.URL, "application/json", bytes.NewReader(data))
}

func createText(feedback *models.Feedback) string {
	emojiTable := []string{
		":grinning:",
		":slightly_smiling_face:",
		":neutral_face:",
		":slightly_frowning_face:",
		":angry:",
	}
	str := "Receive new feedback"
	if len(feedback.Domain) > 0 {
		str += " from <" + feedback.URL + ">"
	}
	if feedback.Emoji >= 1 && feedback.Emoji <= 5 {
		str += " " + emojiTable[feedback.Emoji-1]
	}
	str += "\n"
	str += "> " + feedback.Message
	if len(feedback.Email) > 0 || len(feedback.UserID) > 0 {
		str += "> - "
		if len(feedback.Email) > 0 {
			str += feedback.Email
		} else {
			str += "user-id:"
		}
		if len(feedback.UserID) > 0 {
			str += " " + feedback.UserID
		}
	}
	return str
}
