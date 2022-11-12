package main

import (
	"fmt"
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fmt.Println("Serving on http://localhost:8000")
	getRoutes()
	http.ListenAndServe(":8000", nil)
}

func getRoutes() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request ) {
	templates.ExecuteTemplate(w, "Index", nil )
}