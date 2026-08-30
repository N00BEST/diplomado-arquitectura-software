package main

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/go-sql-driver/mysql"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func loadDBConfig() DBConfig {
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	return DBConfig{
		Host:     host,
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		c.User, c.Password, c.Host, c.Port, c.Database)
}

func openDB(cfg DBConfig) (*sql.DB, error) {
	if cfg.User == "" || cfg.Password == "" || cfg.Database == "" {
		return nil, fmt.Errorf("DB_USER, DB_PASSWORD, and DB_DATABASE are required")
	}

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, err
	}

	var pingErr error
	for attempt := 1; attempt <= 10; attempt++ {
		pingErr = db.Ping()
		if pingErr == nil {
			return db, nil
		}
		log.Printf("waiting for database (attempt %d/10): %v", attempt, pingErr)
		time.Sleep(2 * time.Second)
	}

	db.Close()
	return nil, pingErr
}

func runMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	source, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", source, "mysql", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
