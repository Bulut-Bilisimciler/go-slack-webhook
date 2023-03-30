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

	attachment := models.Attachment{
		Color:   stringPtr("#36a64f"),
		PreText: stringPtr("Warning"),
		Text:    stringPtr("Error occured"),

		MarkdownIn: &[]string{"text", "pretext", "fields"},
	}
	attachment.AddField(models.Field{
		Title: "State",
		Value: "Not urgent",
		Short: true,
	})
	attachment.AddField(models.Field{
		Title: "Message",
		Value: "Mission completed.",
		Short: false,
	})

	attachment.AddAction(models.Action{

		Text:  "Generate error report",
		Type:  "button",
		Style: "danger",
	})
	attachment.AddAction(models.Action{

		Text:  "Fix",
		Type:  "button",
		Style: "danger",
	})

	// create payload
	payload := models.Payload{
		Channel:     channel,
		Username:    ("Hello from go pkg."),
		IconEmoji:   (":robot_face:"),
		Attachments: []models.Attachment{attachment},

		IconUrl:  ":warning:",
		Text:     "Hello from go pkg.",
		Markdown: true,
	}

	// send message
	err := sc.SendWebhook(&payload)
	if err != nil {
		log.Fatal(err)
	}

}
func stringPtr(str string) *string {
	return &str
}
