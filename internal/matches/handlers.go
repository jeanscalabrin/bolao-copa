package matches

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanscalabrin/bolao-copa/internal/shared/render"
)

func ListMatches(c *gin.Context) {
	matches := List()

	render.HTML(c, Page(matches))
}
