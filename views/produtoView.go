package views

import (
	"app/controllers"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request ) {
	items := controllers.BuscaTodosProdutos()
	templates.ExecuteTemplate(w, "Index", items)
}