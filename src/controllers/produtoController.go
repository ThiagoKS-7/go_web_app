package controllers

import (
	"app/src/models"
)

func BuscaTodosProdutos() []models.Produto {
	produtos := models.BuscaTodosProdutos()
	if produtos == nil {
		panic("Nenhum produto encontrado no banco.")
	}
	return produtos
}