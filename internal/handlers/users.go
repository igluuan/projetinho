package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"projetinho/internal/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Buscando usuários")
	idStr := r.URL.Query().Get("id")
	if idStr != "" {
		GetUserByID(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
	log.Println("Usuários retornados com sucesso")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Criando novo usuário")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Erro ao decodificar o corpo da requisição:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = len(models.Users) + 1
	user.CreatedAt = time.Now()
	models.Users = append(models.Users, user)
	log.Println("Usuário criado com sucesso:", user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Buscando usuário por ID")
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	for _, user := range models.Users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			log.Println("Usuário retornado com sucesso:", user)
			return
		}
	}
	http.NotFound(w, r)
}
