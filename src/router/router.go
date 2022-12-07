package router

import (
	"net/http"
	"app/src/views"
)

func GetRoutes() {
	http.HandleFunc("/", views.Index)
	http.HandleFunc("/new", views.New)
}

