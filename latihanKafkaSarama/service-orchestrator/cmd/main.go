package main

import (
	"kafka-sarama-try/internal/provider/routes"

	"github.com/rs/zerolog/log"
)

func main() {
	err := routes.SetupRoutes().Run("127.0.0.1:8081")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

// kafkaConfig := domain.KafkaConfig{
// Brokers: []string{"127.0.0.1:29092"},
// Topic: "Topic_Didi_Test",
// }

// producer, err := kafka.NewKafkaProducer(kafkaConfig.Brokers)
// if err != nil {
// 	log.Fatalf("Failed to create Kafka producer: %v", err)
// }
// defer producer.Close()

// partition, offset, err := kafka.SendMessage(producer, kafkaConfig.Topic, "Hello remon")
// if err != nil {
// 	log.Fatalf("Failed to send message: %v", err)
// }

// fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
