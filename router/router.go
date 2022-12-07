package router

import (
	"net/http"
	"app/views"
)

func GetRoutes() {
	http.HandleFunc("/", views.Index)
}

