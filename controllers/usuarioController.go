package controllers

import (
	"github/arthurgal/Projeto-Eng-de-Software/database"
	"github/arthurgal/Projeto-Eng-de-Software/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriaUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&usuario)
	c.JSON(http.StatusOK, usuario)
}

func ListaUsuario(c *gin.Context) {
	var usuario []models.Usuario

	database.DB.Find(&usuario)
	c.JSON(200, usuario)

}

func EditaUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.First(&usuario, id)

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&usuario).UpdateColumns(usuario)
	c.JSON(http.StatusOK, usuario)
}

func BuscaUsuarioPorID(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.First(&usuario, id)

	if usuario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "usuario não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func DeletaUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.Delete(&usuario, id)

	c.JSON(http.StatusOK, gin.H{
		"Deletado": "usuario foi deletado"})

}
