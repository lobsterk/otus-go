package sender

import (
	"context"
	"encoding/json"

	"github.com/lobsterk/otus-go/hw12_13_14_15_calendar/internal/entities"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Sender struct {
	service entities.EvenetsSeviceInterface
}

func NewSender(handler entities.EvenetsSeviceInterface) Sender {
	return Sender{
		service: handler,
	}
}

func (sender *Sender) Listen(context context.Context) error {
	zap.L().Info("start listen")

	return sender.service.Listen(
		context,
		func(deliveredMessages <-chan amqp.Delivery) {
			zap.L().Info("get events")
			for message := range deliveredMessages {
				messageText := message.Body
				event := &entities.Event{}
				err := json.Unmarshal(messageText, event)
				if err != nil {
					zap.L().Error("parse error", zap.Error(err))
				} else {
					zap.L().Info("event was created")
				}
			}
		},
	)
}

func (sender *Sender) Stop() error {
	err := sender.service.Stop()
	if err != nil {
		return err
	}

	return nil
}
