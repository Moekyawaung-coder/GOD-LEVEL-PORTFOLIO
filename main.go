package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/api/projects", GetProjects)
	http.HandleFunc("/api/login", Login)

	http.ListenAndServe(":8080", nil)
}
