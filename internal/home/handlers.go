package home

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanscalabrin/bolao-copa/internal/shared/render"
)

func Show(c *gin.Context) {
	render.HTML(c, Plage())
}
