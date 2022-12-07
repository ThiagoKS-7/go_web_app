package models

import "https://github.com/ThiagoKS-7/go_web_app/db"

type Produto struct {
	Id 			int
	Nome 		string
	Descricao 	string
	Preco 		float64
	Quantidade 	int
}

func BuscaTodosProdutos() []Produto {
	db := ConectaComDb()
	selectAll, err := db.Query("SELECT * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto {}
	for selectAll.Next() {
		var id, quantidade	int
		var nome,desc		string
		var preco 			float64

		err = selectAll.Scan(&id, &nome, &desc, &preco,&quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = desc
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}