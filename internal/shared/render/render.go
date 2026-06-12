package render

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func HTML(c *gin.Context, component templ.Component) {
	c.Header("Content-Type", "text/html")

	err := component.Render(c.Request.Context(), c.Writer)

	if err != nil {
		c.Status(500)
	}
}
