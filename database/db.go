package database

import (
	"03_RMS/errorHandling"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	RMS *sqlx.DB
)

type SSLMode string

const (
	SSLModeEnable  SSLMode = "enable"
	SSLModeDisable SSLMode = "disable"
)

func ConnectAndMigrate(host, port, databaseName, user, password string, sslMode SSLMode) error {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, databaseName, sslMode)
	DB, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	RMS = DB
	return migrateUp(DB)
}

func migrateUp(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// Tx provides the transaction wrapper
func Tx(fn func(tx *sqlx.Tx) error) error {
	tx, err := RMS.Beginx()
	if err != nil {
		return errorHandling.UnableToBeginTransaction()
	}
	defer func() error {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				return errorHandling.UnableToRollbackTransaction()
			}
		}
		if commitErr := tx.Commit(); commitErr != nil {
			return errorHandling.UnableToCommitTransaction()
		}
		return nil
	}()
	err = fn(tx)
	return err
}
