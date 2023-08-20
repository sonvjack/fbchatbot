package fb

import (
	"fbchatbot/config"

	messenger "github.com/mileusna/facebook-messenger"
)

type FacebookClient struct {
	Client *messenger.Messenger
}

// 258826560297204
// 169e8f3c03c1257eb25aa3151354b166
func NewFacebookClient(cfg *config.Config,
	messageReceived func(msng *messenger.Messenger, userID int64, m messenger.FacebookMessage),
	postbackReceived func(msng *messenger.Messenger, userID int64, p messenger.FacebookPostback),
	deliveryReceived func(msng *messenger.Messenger, userID int64, d messenger.FacebookDelivery),
) *FacebookClient {
	msng := &messenger.Messenger{
		AccessToken:     cfg.AccessToken,
		VerifyToken:     cfg.VerifyToken,
		PageID:          cfg.PageID,
		MessageReceived: messageReceived, // your function for handling received messages, defined below
	}

	// in init or afterwards, you can also specify events when receiving postbacks and message delivery reports from Facebook
	// if you don't want to manage this events, just comment/don't use this receivers
	msng.PostbackReceived = postbackReceived // comment/delete if not used
	msng.DeliveryReceived = deliveryReceived // comment/delete if not used

	// set URL for your webhook and directly use msng as http Handler
	return &FacebookClient{
		Client: msng,
	}
}
