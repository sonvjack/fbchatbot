package handler

import (
	messenger "github.com/mileusna/facebook-messenger"
)

// PostbackReceived is called when you reiceive postback event from Facebook server
func PostbackReceived(msng *messenger.Messenger, userID int64, p messenger.FacebookPostback) {
	if p.Payload == "THIS_DATA_YOU_WILL_RECEIVE_AS_POSTBACK_WHEN_USER_CLICK_THE_BUTTON" {
		// user just clicked Ok button from previouse example, lets just send him a message
		MessageChan <- MsgTemplate{
			UserID:  userID,
			Msng:    msng,
			Message: "Ok, I'm always online, chat with me anytime :)",
		}
	}
}
