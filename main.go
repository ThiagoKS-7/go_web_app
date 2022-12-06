package main

import (
	"fmt"
	"net/http"
	"html/template"
	_ "github.com/lib/pq"
	_ "github.com/joho/godotenv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome 		string
	Descricao 	string
	Preco 		float64
	Quantidade 	int
}

type Conexao struct {
	DB_Name 		string
	DB_User 		string
	DB_Password 	string
	DB_Host 		string
	DB_Ssl 			string
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}

func conectaComDb() *sql.DB {
	env := {
		DB_Name: 	 	goDotEnvVariable("DB_NAME"),
		DB_User: 	 	goDotEnvVariable("DB_USER"),
		DB_Password:	goDotEnvVariable("DB_PASSWORD"),
		DB_Host 		goDotEnvVariable("DB_HOST")
		DB_Ssl 			goDotEnvVariable("DB_SSL")
	}
	dbname := 
	dbuser := goDotEnvVariable("DB_USER")
	dbpwd := goDotEnvVariable("DB_PASSWORD")


	con := "user dbname password host sslmode"
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