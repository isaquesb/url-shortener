package rabbitmq

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

func NewSubscriber(logger watermill.LoggerAdapter, consumerGroup string, rabbitURL string) (message.Subscriber, error) {
	config := amqp.NewDurablePubSubConfig(
		rabbitURL,
		amqp.GenerateQueueNameTopicName,
	)
	config.Exchange.Type = "direct"
	subscriber, err := amqp.NewSubscriber(config, logger)
	if err != nil {
		return nil, err
	}
	return subscriber, nil
}
