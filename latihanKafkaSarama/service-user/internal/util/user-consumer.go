package util

import (
	"context"
	"log"
	"service-user-kafka/internal/domain"
	"service-user-kafka/internal/kafka"
	"service-user-kafka/internal/usecase"
	"sync"
)

// for consuming the message that were sent to topic 'user_topic_validate'
func StartConsumeFromOrder(usecase *usecase.UserUsecase) {
	kafkaConfig := domain.KafkaConfig{
		Brokers: []string{"127.0.0.1:29092"}, // Replace with your Kafka broker addresses
		GroupID: "user-consumer-group",
		Topic:   "user_topic_validate",
	}

	// setup kafka producer for kafka for user and package
	producer, err := kafka.NewKafkaProducer(kafkaConfig.Brokers)
	if err != nil {
		log.Printf("Failed to produce new producer for kafka: %s\n", err)
		return
	}
	defer producer.Close()

	handler := kafka.NewMessageHandler(producer, usecase)

	consumer, err := kafka.NewKafkaConsumer(kafkaConfig.Brokers, kafkaConfig.GroupID, []string{kafkaConfig.Topic}, handler)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := consumer.Consume(ctx); err != nil {
			log.Printf("Error from consumer: %v", err)
		}
	}()
	wg.Wait()
}
