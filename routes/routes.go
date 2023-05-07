package routes

import (
	"github.com/gin-gonic/gin"
	"github/arthurgal/Projeto-Eng-de-Software/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.POST("/produto", controllers.CriaProduto)
	r.GET("/produto", controllers.ListaProduto)
	r.GET("/produto/:id", controllers.BuscaProdutoPorID)
	r.PUT("/produto/:id", controllers.EditaProduto)
	r.DELETE("/produto/:id", controllers.DeletaProduto)
	r.Run()
}