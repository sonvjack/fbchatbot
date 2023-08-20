package handler

import (
	"fbchatbot/logger"

	"go.uber.org/zap"

	messenger "github.com/mileusna/facebook-messenger"
)

// DeliveryReceived is used if you want to track delivery reports for sent messages
func DeliveryReceived(msng *messenger.Messenger, userID int64, d messenger.FacebookDelivery) {
	for _, mid := range d.Mids {
		logger.Logger.Info("Message delivered", zap.String("msgID", mid))
	}
}
