package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jacob/modules/modules/bootstrap/config"
)

func Middleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Middleware logic before request
		log.Printf("config: %s", config)
		c.Next()
		// Middleware logic after request
	}
}
