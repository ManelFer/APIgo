package routes

import (
	"github.com/ManelFer/APIgo/in/handlers"
	"github.com/ManelFer/APIgo/in/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Rotas públicas (sem JWT)
	r.POST("/users", handlers.CreateUser)
	r.POST("/login", handlers.Login)

	// Rotas protegidas (exigem Authorization: Bearer <token>)
	protected := r.Group("")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/users", handlers.Getusers)
	}
}
