package handler

import (
	messenger "github.com/mileusna/facebook-messenger"
)

// MessageReceived is called when you receive message on you webhook i.e. when someone sends message to your chat bot
// params: messenger that received the message, then the user id that sent us message and message data itself
func MessageReceived(msng *messenger.Messenger, userID int64, m messenger.FacebookMessage) {

	// message received, now lets check what user has sent to us
	switch m.Text {
	// can through mysql or other database set more reply words
	case "hello", "hi":
		// someone sent hello or hi, reply with simple text message
		MessageChan <- MsgTemplate{
			UserID:  userID,
			Msng:    msng,
			Message: "Hello there",
		}

	case "send me website":
		// now lets send him some structured message with image and link
		gm := msng.NewGenericMessage(userID)
		gm.AddNewElement("Title", "Subtitle", "http://mysite.com", "http://mysite.com/some-photo.jpeg", nil)

		// GenericMessage can contain up to 10 elements, they are represented as cards and can be scoreled horicontally in messenger
		// So lets add one more element, this time with buttons
		btn1 := msng.NewWebURLButton("Contact US", "http://mysite.com/contact")
		btn2 := msng.NewPostbackButton("Ok", "THIS_DATA_YOU_WILL_RECEIVE_AS_POSTBACK_WHEN_USER_CLICK_THE_BUTTON")
		gm.AddNewElement("Site title", "Subtitle", "http://mysite.com", "http://mysite.com/some-photo.jpeg", []messenger.Button{btn1, btn2})

		// ok, message is ready, lets send
		MessageChan <- MsgTemplate{
			UserID:  userID,
			Msng:    msng,
			Message: gm,
		}

	default:
		// upthere we haven't check for errors and responses for cleaner example code
		// but keep in mind that SendMessage returns FacebookResponse struct and error
		// errors are received from Facebook if sometnihg went wrong with message sending
		MessageChan <- MsgTemplate{
			UserID:  userID,
			Msng:    msng,
			Message: m.Text,
		}
	}
}
