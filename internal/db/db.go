package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/melinaco4/grpc-service/internal/ticket"
)

type DB struct {
	db *sqlx.DB
}

func New() (DB, error) {
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
		return DB{}, err
	}
	return db, nil
}

func (d DB) GetTicketByID(id string) (ticket.Ticket, error) {
	return ticket.Ticket{}, nil
}

func (d DB) InsertTicket(t ticket.Ticket) (ticket.Ticket, error) {
	return ticket.Ticket{}, nil
}

func (d DB) DeleteRocket(id string) error {
	return nil
}
