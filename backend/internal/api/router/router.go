package router

import (
	"bounty_board/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	r := gin.New()

	r.GET("/health", h.HealthCheck)

	return r
}
