package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := loadDBConfig()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := runMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	r := gin.Default()
	registerRecordRoutes(r, newRecordStore(db))

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
