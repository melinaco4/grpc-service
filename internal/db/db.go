package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/melinaco4/grpc-service/internal/ticket"
)

type Service struct {
	db *sqlx.DB
}

func New() (Service, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Service{}, err
	}
	return Service{
		db: db,
	}, nil
}

func (s Service) GetTicketByID(id string) (ticket.Ticket, error) {
	return ticket.Ticket{}, nil
}

func (s Service) InsertTicket(t ticket.Ticket) (ticket.Ticket, error) {
	return ticket.Ticket{}, nil
}

func (s Service) DeleteRocket(id string) error {
	return nil
}
