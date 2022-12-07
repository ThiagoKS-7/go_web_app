package db

import (
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

type Conexao struct {
	Name 		string
	User 		string
	Password 	string
	Host 		string
	Ssl 		string
}

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  panic("Error loading .env file")
	}
  
	return os.Getenv(key)
}

func ConectaComDb() *sql.DB {
	env := Conexao {
		Name: 	 	GoDotEnvVariable("DB_NAME"),
		User: 	 	GoDotEnvVariable("DB_USER"),
		Password:	GoDotEnvVariable("DB_PASSWORD"),
		Host: 		GoDotEnvVariable("DB_HOST"),
		Ssl: 		GoDotEnvVariable("DB_SSL"),
	}

	con := "user="+env.User + " dbname="+env.Name + " password="+env.Password + " host="+env.Host + " sslmode=" + env.Ssl
	db,err := sql.Open("postgres",con) 
	if err != nil {
		panic(err.Error())
	}
	return db
}