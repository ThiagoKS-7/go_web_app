package main

import (
	"fmt"
	"net/http"
	"os"
	_db "db"
	"database/sql"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))


func main() {
	fmt.Println("Serving on http://localhost:" + db.GoDotEnvVariable("DB_APP_PORT"))
	getRoutes()
	http.ListenAndServe(":" + db.GoDotEnvVariable("DB_APP_PORT"), nil)
}

func getRoutes() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request ) {
	templates.ExecuteTemplate(w, "Index", produtos)
}