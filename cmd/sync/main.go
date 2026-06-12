package main

import (
	"fmt"
	"log"

	"github.com/jeanscalabrin/bolao-copa/internal/config"
	"github.com/jeanscalabrin/bolao-copa/internal/footballdata"
)

func main() {
	cfg := config.Load()

	client := footballdata.NewClient(
		cfg.FootballDataAPIKey,
	)

	matches, err := client.GetWorldCupMatches()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"matches found: %d\n",
		len(matches),
	)

	fmt.Printf(
		"%+v\n",
		matches[0],
	)
}
