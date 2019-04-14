package messagebird

import (
	"log"

	"github.com/dstotijn/festchat/sms"
	mb "github.com/messagebird/go-rest-api"
	mbsms "github.com/messagebird/go-rest-api/sms"
)

// Client represents a MessageBird API client.
type Client struct {
	c *mb.Client
}

// NewClient returns a new Client.
func NewClient(accessKey string) *Client {
	return &Client{
		c: mb.New(accessKey),
	}
}

// Send sends an SMS message via MessageBird.
func (client Client) Send(msg sms.Message) error {
	sentMsg, err := mbsms.Create(
		client.c,
		msg.Originator,
		msg.Recipients,
		msg.Payload,
		&mbsms.Params{
			DataCoding: "auto",
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Debug: Sent message to MessageBird: %v", sentMsg)

	return nil
}
