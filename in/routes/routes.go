package routes

import (
	"github.com/ManelFer/APIgo/in/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", handlers.Getusers)
}
