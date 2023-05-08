package routes

import (
	"github/arthurgal/Projeto-Eng-de-Software/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	//Rotas para PRODUTO
	r.POST("/produto", controllers.CriaProduto)
	r.GET("/produto", controllers.ListaProduto)
	r.GET("/produto/:id", controllers.BuscaProdutoPorID)
	r.PUT("/produto/:id", controllers.EditaProduto)
	r.DELETE("/produto/:id", controllers.DeletaProduto)

	//Rotas para CARRINHO
	r.POST("/carrinho", controllers.CriaCarrinho)
	r.GET("/carrinho", controllers.ListaCarrinho)
	r.GET("/carrinho/:id", controllers.ListaCarrinhoPorId)
	r.PUT("/carrinho/:id", controllers.EditaCarrinho)
	r.DELETE("/carrinho/:id", controllers.DeletaCarrinho)

	//r.POST("/carrinho/:id/produto", controllers.AdicionaProdutoNoCarrinho)
	r.POST("/carrinho/:id/produto/:produtoID", controllers.CadastraProdutoJaExisteNoCarrinho)

	//Rotas para USUARIO
	r.POST("/usuario", controllers.CriaUsuario)
	r.GET("/usuario", controllers.ListaUsuario)
	r.GET("/usuario/:id", controllers.BuscaUsuarioPorID)
	r.PUT("/usuario/:id", controllers.EditaUsuario)
	r.DELETE("/usuario/:id", controllers.DeletaUsuario)

	r.Run()
}
