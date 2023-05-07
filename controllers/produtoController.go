package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github/arthurgal/Projeto-Eng-de-Software/models"
	"github/arthurgal/Projeto-Eng-de-Software/database"
)

func CriaProduto(c *gin.Context)  {
	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error()})
		return
	}
	database.DB.Create(&produto)
	c.JSON(http.StatusOK, produto)
}

func ListaProduto(c *gin.Context) {
	var produtos []models.Produto
	
	database.DB.Find(&produtos)
	c.JSON(200, produtos)

}


func EditaProduto(c *gin.Context) {
	var produto models.Produto
	id := c.Params.ByName("id")
	database.DB.First(&produto, id)

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&produto).UpdateColumns(produto)
	c.JSON(http.StatusOK, produto)
}

func BuscaProdutoPorID(c *gin.Context) {
	var produto models.Produto
	id := c.Params.ByName("id")
	database.DB.First(&produto, id)

	if produto.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, produto)
}

func DeletaProduto(c *gin.Context){
	var produto models.Produto
	id := c.Params.ByName("id")
	database.DB.Delete(&produto, id)
	
	c.JSON(http.StatusOK, gin.H{
		"Deletado": "produto foi deletado"})
	

}