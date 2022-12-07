package main

import (
	"fmt"
	"net/http"
	"os"
	"database/sql"
	"html/template"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Id 			int
	Nome 		string
	Descricao 	string
	Preco 		float64
	Quantidade 	int
}

type Conexao struct {
	Name 		string
	User 		string
	Password 	string
	Host 		string
	Ssl 		string
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  panic("Error loading .env file")
	}
  
	return os.Getenv(key)
}

func conectaComDb() *sql.DB {
	env := Conexao {
		Name: 	 	goDotEnvVariable("DB_NAME"),
		User: 	 	goDotEnvVariable("DB_USER"),
		Password:	goDotEnvVariable("DB_PASSWORD"),
		Host: 		goDotEnvVariable("DB_HOST"),
		Ssl: 		goDotEnvVariable("DB_SSL"),
	}

	con := "user="+env.User + " dbname="+env.Name + " password="+env.Password + " host="+env.Host + " sslmode=" + env.Ssl
	db,err := sql.Open("postgres",con) 
	if err != nil {
		panic(err.Error())
	}
	return db
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
	db := conectaComDb()
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
	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}