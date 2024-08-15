package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	provider "service-user-kafka/internal/provider/db"
	"service-user-kafka/internal/repository"
	"service-user-kafka/internal/usecase"
	"service-user-kafka/internal/util"
	"syscall"
)

func main() {
	database, err := provider.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepo := repository.NewUserRepo(database)
	userUsecase := usecase.NewUserUsecase(userRepo)

	util.StartConsumeFromOrder(userUsecase)
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan

	log.Println("Received termination signal. Initiating shutdown...")
	cancel()

	log.Println("Consumer has shut down")
}
