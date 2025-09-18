package main

import (
	"auth/config"
	"auth/internal/storage/postgres"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	cfg := config.MustLoadConfig()
	storage, err := postgres.New(cfg.DSN)
	if err != nil {
		return
	}
	if err := goose.Up(storage.DB, cfg.MigrationsPath); err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}
	fmt.Println("migrations applied successfully")
}
