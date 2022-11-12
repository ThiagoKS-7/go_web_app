package main

import (
	"fmt"
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome 		string
	Descricao 	string
	Preco 		float64
	Quantidade 	int
}

func main() {
	fmt.Println("Serving on http://localhost:8000")
	getRoutes()
	http.ListenAndServe(":8000", nil)
}

func getRoutes() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request ) {
	produtos := []Produto{
		{ Nome:"Camiseta", Descricao:"Teste 123", Preco: 39,Quantidade: 5 },
		{ Nome:"Tênis", Descricao:"Teste 123", Preco: 299, Quantidade: 3},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}