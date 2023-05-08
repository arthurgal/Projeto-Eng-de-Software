package models

import (
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Nome       string  `json:"nome"`
	Valor      float64 `json:"valor"`
	Descricao  string  `json:"descricao"`
	Qtd        int     `json:"qtd"`
	CarrinhoID uint
}
