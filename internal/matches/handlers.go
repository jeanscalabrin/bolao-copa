package matches

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeanscalabrin/bolao-copa/internal/shared/render"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(c *gin.Context) {
	matches, err := h.service.List(c.Request.Context())

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	render.HTML(c, Page(matches))
}
