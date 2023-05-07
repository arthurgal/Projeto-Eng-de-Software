package main

import (
	"github/arthurgal/Projeto-Eng-de-Software/routes"
	"github/arthurgal/Projeto-Eng-de-Software/database"

)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}