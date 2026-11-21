package middleware

import (
	"bounty_board/internal/api/apiconsts"
	"context"
	"log/slog"
	"net/http"
	"uuid"

	"time"

	"github.com/gin-gonic/gin"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()

		ctx := context.WithValue(
			c.Request.Context(),
			apiconsts.RequestIDKey,
			requestID,
		)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func LoggingMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		requestID := GetRequestID(c.Request.Context())
		requestLogger := logger.With(
			"request_id", requestID,
		)

		ctx := context.WithValue(c.Request.Context(), apiconsts.RequestLogger, requestLogger)
		c.Request = c.Request.WithContext(ctx)
		c.Next()

		requestLogger.Info(
			"request completed",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", time.Since(start),
		)
	}
}

func RecoverMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				requestID := GetRequestID(c.Request.Context())
				logger.Error(
					"panic recovered",
					"request_id", requestID,
					"error", err,
				)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
				return
			}
		}()
	}
}

func TimeoutMiddleware(timeout int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(timeout)*time.Second)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func RateLimitingMiddleware() {

}

func GetRequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(apiconsts.RequestIDKey).(string)

	return requestID
}
