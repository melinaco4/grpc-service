package db

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (db *DB) Migrate() error {
	driver, err := postgres.WithInstance(db.db.DB, &postgres.COnfig{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
}
