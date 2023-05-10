package database

import (
	"log"

	"github/arthurgal/Projeto-Eng-de-Software/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=postgres password=postgres dbname=root port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Produto{})
	DB.AutoMigrate(&models.Carrinho{})
	DB.AutoMigrate(&models.Usuario{})
}
