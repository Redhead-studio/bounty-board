package handler

import (
	"bounty_board/internal/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	app *application.Application
}

func NewHandler(app *application.Application) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "OK",
	})
}
