package controllers

import (
	"app/models"
)

func BuscaTodosProdutos() []models.Produto {
	produtos := models.BuscaTodosProdutos()
	if produtos == nil {
		panic("Nenhum produto encontrado no banco.")
	}
	return produtos
}