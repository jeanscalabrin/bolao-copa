package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanscalabrin/bolao-copa/internal/home"
	"github.com/jeanscalabrin/bolao-copa/internal/matches"
)

func Setup() *gin.Engine {
	r := gin.Default()

	home.RegisterRoutes(r.Group(""))
	matches.RegisterRoutes(r.Group("/matches"))
	return r
}
