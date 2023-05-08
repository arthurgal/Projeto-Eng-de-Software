package controllers

import (
	"github/arthurgal/Projeto-Eng-de-Software/database"
	"github/arthurgal/Projeto-Eng-de-Software/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriaCarrinho(c *gin.Context) {
	var carrinho models.Carrinho
	if err := c.ShouldBindJSON(&carrinho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&carrinho)
	c.JSON(http.StatusOK, carrinho)
}

func ListaCarrinho(c *gin.Context) {
	var carrinhos []models.Carrinho

	database.DB.Preload("Produto").Find(&carrinhos)
	c.JSON(200, carrinhos)

}

func EditaCarrinho(c *gin.Context) {
	var carrinho models.Carrinho
	id := c.Params.ByName("id")
	database.DB.First(&carrinho, id)

	if err := c.ShouldBindJSON(&carrinho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&carrinho).UpdateColumns(carrinho)
	c.JSON(http.StatusOK, carrinho)
}

func DeletaCarrinho(c *gin.Context) {
	var carrinho models.Carrinho
	id := c.Params.ByName("id")
	database.DB.Delete(&carrinho, id)

	c.JSON(http.StatusOK, gin.H{
		"Deletado": "carrinho foi deletado"})

}

/* func AdicionaProdutoNoCarrinho(c *gin.Context) {
	id := c.Param("id")

	var carrinho models.Carrinho
	if err := database.DB.First(&carrinho, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Carrinho não encontrado"})
		return
	}

	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	produto.CarrinhoID = carrinho.ID
	if err := database.DB.Create(&produto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	c.Status(http.StatusCreated)
}
*/

func ListaCarrinhoPorId(c *gin.Context) {
	id := c.Param("id")

	var carrinho models.Carrinho
	if err := database.DB.Preload("Produto").First(&carrinho, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter carrinho"})
		return
	}

	c.JSON(http.StatusOK, carrinho)
}

func CadastraProdutoJaExisteNoCarrinho(c *gin.Context) {
	id := c.Param("id")
	produtoID := c.Param("produtoID")

	var carrinho models.Carrinho
	if err := database.DB.First(&carrinho, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Carrinho não encontrado"})
		return
	}

	var produto models.Produto
	if err := database.DB.First(&produto, produtoID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Produto não encontrado"})
		return
	}

	// Associar o produto ao carrinho
	produto.CarrinhoID = carrinho.ID
	if err := database.DB.Save(&produto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao associar o produto ao carrinho"})
		return
	}

	c.Status(http.StatusOK)
}
