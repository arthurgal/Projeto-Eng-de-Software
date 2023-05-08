package models

import (
	"gorm.io/gorm"
)

type Carrinho struct {
	gorm.Model
	Produto []Produto `gorm:"foreignKey:CarrinhoID"`
}
