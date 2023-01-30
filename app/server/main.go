package main

import (
	"log"

	"github.com/melinaco4/grpc-service/internal/db"
	"github.com/melinaco4/grpc-service/internal/ticket"
)

func Run() error {
	// this will be responsible for initializing and starting grpc server
	ticketStore, err := db.New()
	if err != nil {
		return err
	}
	err = ticketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	_ = ticket.New(ticketStore)
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
