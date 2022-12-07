package views

import (
	"app/src/controllers"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("src/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request ) {
	items := controllers.BuscaTodosProdutos()
	templates.ExecuteTemplate(w, "Index", items)
}

func New(w http.ResponseWriter, r *http.Request ) {
	//items := controllers.BuscaTodosProdutos()
	templates.ExecuteTemplate(w, "New", r)
}