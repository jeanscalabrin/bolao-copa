package main

import (
	"context"
	"log"

	"github.com/jeanscalabrin/bolao-copa/internal/config"
	"github.com/jeanscalabrin/bolao-copa/internal/database"
	"github.com/jeanscalabrin/bolao-copa/internal/footballdata"
	"github.com/jeanscalabrin/bolao-copa/internal/matches"
	syncservice "github.com/jeanscalabrin/bolao-copa/internal/sync"
)

func main() {
	cfg := config.Load()

	db, err := database.New(
		cfg.DatabaseURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	client := footballdata.NewClient(
		cfg.FootballDataAPIKey,
	)

	repo := matches.NewRepository(db)

	service := syncservice.NewService(
		client,
		repo,
	)

	err = service.SyncMatches(
		context.Background(),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("matches synced")
}
