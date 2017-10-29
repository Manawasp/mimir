package config

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"

	log "github.com/Sirupsen/logrus"

	"mimir/api/models"
)

// WebhookChannel TODO
type WebhookChannel struct {
	Name    string `toml:"channel"`
	Pattern string `toml:"pattern"`
}

// Webhook TODO
type Webhook struct {
	URL      string           `toml:"url"`
	Icon     string           `toml:"icon"`
	Username string           `toml:"username"`
	Channels []WebhookChannel `toml:"channel"`
}

// NotifyMessage TODO
type NotifyMessage struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
	Icon     string `json:"icon_url"`
}

// Notify TODO
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
	data, _ := json.Marshal(notif)
	_, err := http.Post(wh.URL, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Errorf("Error: %v.", err)
	}
}

func createText(feedback *models.Feedback) string {
	emojiTable := []string{
		":angry:",
		":slightly_frowning_face:",
		":neutral_face:",
		":slightly_smiling_face:",
		":grinning:",
	}
	str := "Receive new feedback"
	if len(feedback.Domain) > 0 {
		str += " from <" + feedback.Domain + ">"
	}
	str += ":\n"
	str += "> " + feedback.Message
	str += "\n"
	if len(feedback.Email) > 0 || len(feedback.UserID) > 0 {
		if len(feedback.Email) > 0 {
			str += "*email*: " + feedback.Email
		}
		if len(feedback.UserID) > 0 {
			str += "  *id*: #" + feedback.UserID
		}
	}
	if feedback.Emoji >= 1 && feedback.Emoji <= 5 {
		str += "  " + emojiTable[feedback.Emoji-1]
	}
	return str
}
