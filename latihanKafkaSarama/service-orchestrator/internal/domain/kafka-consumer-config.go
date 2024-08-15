package domain

import "github.com/IBM/sarama"

type KafkaConsumer struct {
	Consumer sarama.ConsumerGroup
	Handler  sarama.ConsumerGroupHandler
	Topics   []string
}
