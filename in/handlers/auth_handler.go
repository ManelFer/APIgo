// Login e rotinas de autenticação
package handlers

import (
	"net/http"

	"github.com/ManelFer/APIgo/in/auth"
	"github.com/ManelFer/APIgo/in/models"
	"github.com/ManelFer/APIgo/in/repositories"
	"github.com/gin-gonic/gin"
)

// LoginRequest corpo esperado do POST /login
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login verifica email/senha e retorna um JWT.
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email e senha são obrigatórios"})
		return
	}

	user, err := repositories.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	if !auth.CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	token, err := auth.Generatetoken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  toPublicUser(user),
	})
}

// toPublicUser remove a senha da resposta
func toPublicUser(u *models.User) gin.H {
	return gin.H{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
	}
}
