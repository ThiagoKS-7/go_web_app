package main

import (
	"fmt"
	"net/http"
	"app/src/router"
	"app/database"
)

func main() {
	fmt.Println("Serving on http://localhost:" + db.GoDotEnvVariable("DB_APP_PORT"))
	router.GetRoutes()
	http.ListenAndServe(":" + db.GoDotEnvVariable("DB_APP_PORT"), nil)
}

