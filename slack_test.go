package slackwebhook

import (
	"log"
	"testing"

	"github.com/Bulut-Bilisimciler/go-slack-webhook/models"
)

func TestSlackSendWebhook(t *testing.T) {
	// read config
	webhookURL := "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX"
	channel := "#test"

	// init slack config
	sc := InitSlackWebhookConfig(webhookURL, channel)

	// create payload
	payload := models.Payload{
		Channel:     channel,
		Username:    ("Hello from go pkg."),
		IconEmoji:   (":robot_face:"),
		Attachments: []models.Attachment{},
		Text:        "Hello from go pkg.",
	}

	// send message
	err := sc.SendWebhook(&payload)
	if err != nil {
		log.Fatal(err)
	}
}
