package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jeanscalabrin/bolao-copa/internal/config"
	"github.com/jeanscalabrin/bolao-copa/internal/database"
	"github.com/jeanscalabrin/bolao-copa/internal/home"
	"github.com/jeanscalabrin/bolao-copa/internal/matches"
)

func Setup() *gin.Engine {
	cfg := config.Load()

	db, err := database.New(cfg.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	repo := matches.NewRepository(db)
	service := matches.NewService(repo)
	handler := matches.NewHandler(service)

	home.RegisterRoutes(r.Group(""))
	matches.RegisterRoutes(r.Group("/matches"), handler)
	return r
}
