package main

import (
	"net/http"

	"github.com/ManelFer/APIgo/in/config/database"
	"github.com/ManelFer/APIgo/in/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// r = router
	r := gin.Default()
	database.ConnectarBancoDados() // Conectando ao banco de dados
	routes.RegisterRoutes(r)       // Registrando as rotas

	// definindo um simples get endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// iniciando o servidor na porta 8080

	r.Run(":8080")
}
