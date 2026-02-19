// simualando para teste
package handlers

import (
	"net/http"

	"github.com/ManelFer/APIgo/in/auth"
	"github.com/ManelFer/APIgo/in/models"
	"github.com/ManelFer/APIgo/in/repositories"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	// sempre manter a leitura do JSON antes de qualquer outra coisa
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos",
		})
		return
	}

	// validar os dados do usuário
	if user.Email == "" || user.Password == "" || user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nome, email e senha são obrigatórios",
		})
		return
	}
	// criptografar a senha
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criptografar a senha",
		})
		return
	}
	user.Password = hashedPassword

	//salvar o usuário no banco de dados
	err = repositories.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar usuário",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "usuário criado com sucesso",
	})
}

func Getusers(c *gin.Context) {
	users, err := repositories.GetAllusers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter usuários"})
		return
	}

	c.JSON(http.StatusOK, users)
}
