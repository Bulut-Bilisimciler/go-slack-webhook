## Go-Slack-Webhook  
This is the simple repository for sending slack message with golang.

## Installation

```bash
go get github.com/Bulut-Bilisimciler/go-slack-webhook
```

### Example

- Example golang code

```go
func main() {
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
```


### Change Log

- v1.0.0
  - Initial release (simple payload and message send utility added)