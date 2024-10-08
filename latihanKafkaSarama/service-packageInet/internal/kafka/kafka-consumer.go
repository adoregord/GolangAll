package kafka

import (
	"buy-package-kafka/internal/domain"
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	consumer sarama.ConsumerGroup
	handler  sarama.ConsumerGroupHandler
	topics   []string
}

type MessageHandler struct {
	Producer sarama.SyncProducer
}

func NewMessageHandler(producer sarama.SyncProducer) *MessageHandler {
	return &MessageHandler{
		Producer: producer,
	}
}

func (h MessageHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h MessageHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h MessageHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	if h.Producer == nil {
		log.Println("Producer is nil! Make sure it's initialized properly.")
		return errors.New("kafka producer is not initialized")
	}

	for msg := range claim.Messages() {
		sess.MarkMessage(msg, "")

		// Parse the incoming message
		var incoming domain.IncomingMessage
		if err := json.Unmarshal(msg.Value, &incoming); err != nil {
			log.Printf("Error parsing message: %v\n", err)
			continue
		}

		switch incoming.OrderType {
		case "Buy Package":
			// set the service so orches know where the message came from
			incoming.OrderService = "activatePackage"

			log.Println("Adding Paket Internet GG to user: ", incoming.UserId)

			responseBytes, err := json.Marshal(incoming)
			if err != nil {
				log.Printf("Failed to marshal message: %v\n\n", err)
				continue
			}

			_, _, err = h.Producer.SendMessage(&sarama.ProducerMessage{
				Topic: "topic_order",
				Key:   sarama.ByteEncoder(msg.Key),
				Value: sarama.ByteEncoder(responseBytes),
			})
			if err != nil {
				log.Printf("Error writing message to topic_validateUser: %v\n\n", err)
				continue
			}
			log.Printf("Message sent to topic_order: %s\n\n", string(responseBytes))

		default:
			log.Printf("Received unsupported message format: %v\n\n", incoming)
		}
	}
	return nil
}

func NewKafkaConsumer(brokers []string, groupID string, topics []string, msg *MessageHandler) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer: consumer,
		topics:   topics,
		handler:  msg,
	}, nil
}

func (kc *KafkaConsumer) Consume(ctx context.Context) error {
	for {
		err := kc.consumer.Consume(ctx, kc.topics, kc.handler)
		if err != nil {
			return err
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (kc *KafkaConsumer) Close() error {
	return kc.consumer.Close()
}

func (kc *KafkaConsumer) SetHandler(handler sarama.ConsumerGroupHandler) {
	kc.handler = handler
}
