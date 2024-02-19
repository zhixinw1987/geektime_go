package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func EnableCorsMiddleware(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		MaxAge:           time.Duration(5) * time.Minute,
	}))
}
