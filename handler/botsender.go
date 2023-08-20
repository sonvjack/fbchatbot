package handler

import (
	"context"
	"fbchatbot/logger"
	"sync"

	"go.uber.org/zap"

	messenger "github.com/mileusna/facebook-messenger"
)

type MsgTemplate struct {
	Msng    *messenger.Messenger
	Message interface{}
	UserID  int64
}

var MessageChan chan MsgTemplate

func SendFBMsg(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	logger.Logger.Info("send message service start")
	for {
		select {
		case <-ctx.Done():
			logger.Logger.Info("send message service stop")
			return
		case msg := <-MessageChan:
			switch msg.Message.(type) {
			case string:
				resp, err := msg.Msng.SendTextMessage(msg.UserID, msg.Message.(string))
				if err != nil {
					logger.Logger.Error("SendTextMessage Error", zap.Error(err),
						zap.Int64("send to user", msg.UserID))
					return // if there is an error, resp is empty struct, useless
				}
				logger.Logger.Info("Message Received", zap.String("Message ID", resp.MessageID),
					zap.Int64("sent to user", resp.RecipientID))
			case messenger.GenericMessage:
				resp, err := msg.Msng.SendMessage(msg.Message.(messenger.GenericMessage))
				if err != nil {
					logger.Logger.Error("SendTextMessage Error", zap.Error(err))
					return // if there is an error, resp is empty struct, useless
				}
				logger.Logger.Info("Message Received", zap.String("Message ID", resp.MessageID),
					zap.Int64("sent to user", resp.RecipientID))
			default:
				logger.Logger.Error("unknown message type", zap.Any("msg", msg.Message))
			}
		}
	}
}
