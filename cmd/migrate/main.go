package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jeanscalabrin/bolao-copa/internal/config"
)

func main() {
	cfg := config.Load()

	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	log.Println("migrations concluded")
}
