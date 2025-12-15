package main

import (
	"log"
	"net/http"

	"projetinho/internal/handlers"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetUsers(w, r)
	case "POST":
		handlers.CreateUser(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
