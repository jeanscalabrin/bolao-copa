package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanscalabrin/bolao-copa/internal/home"
)

func Setup() *gin.Engine {
	r := gin.Default()

	home.RegisterRoutes(r.Group(""))
	return r
}
