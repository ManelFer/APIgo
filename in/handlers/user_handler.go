// simualando para teste
package handlers

import (
	"net/http"

	"github.com/ManelFer/APIgo/in/repositories"
	"github.com/gin-gonic/gin"
)

func Getusers(c *gin.Context) {
	users, err := repositories.GetAllusers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter usuários"})
		return
	}

	c.JSON(http.StatusOK, users)
}
