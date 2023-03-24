package slackwebhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-slack-webhook/models"
)

// SlackConfig struct
type SlackConfig struct {
	WebhookURL string `yaml:"webhook_url"`
	Channel    string `yaml:"channel"`
}

func InitSlackWebhookConfig(webhookUrl string, channel string) *SlackConfig {
	return &SlackConfig{
		WebhookURL: webhookUrl,
		Channel:    channel,
	}
}

func (sc *SlackConfig) SendWebhook(payload *models.Payload) error {
	// fill the channel from config
	payload.Channel = sc.Channel

	// try to marshall the payload
	body, err := json.Marshal(payload)
	if err != nil {
		return errors.New("unable to jsonize payload" + err.Error())
	}

	// fire the webhook
	resp, err := http.Post(sc.WebhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return errors.New("unable to post payload to slack " + err.Error())
	}
	defer resp.Body.Close()

	// return success
	return nil
}

func SlackString(str string) *string {
	return &str
}
