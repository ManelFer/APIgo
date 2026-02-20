package handlers

import (
	"net/http"

	"github.com/ManelFer/APIgo/in/models"
	"github.com/ManelFer/APIgo/in/repositories"
	"github.com/gin-gonic/gin"
)

// CreateEstoque insere um novo item no estoque (requer login).
func CreateEstoque(c *gin.Context) {
	var e models.Estoque
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	if e.Equipamento == "" || e.Marca == "" || e.Modelo == "" || e.Patrimonio == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipamento, marca, modelo e patrimonio são obrigatórios"})
		return
	}
	if e.Quantidade < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantidade não pode ser negativa"})
		return
	}
	if err := repositories.CreateEstoque(e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir no estoque"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Item inserido no estoque com sucesso"})
}

// GetEstoque lista todos os itens do estoque (requer login).
func GetEstoque(c *gin.Context) {
	list, err := repositories.GetAllEstoque()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter estoque"})
		return
	}
	c.JSON(http.StatusOK, list)
}
