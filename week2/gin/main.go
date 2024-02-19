package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhixinw1987/geektime_go/week2/internal/middleware"
	web "github.com/zhixinw1987/geektime_go/week2/internal/web/user"
)

func main() {
	server := gin.Default()
	handler := &web.UserHandler{}
	handler.RegisterUserRoutes(server)

	middleware.EnableCorsMiddleware(server)
	server.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
