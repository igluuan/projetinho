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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = len(models.Users) + 1
	user.CreatedAT = time.Now()
	models.Users = append(models.Users, user)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	for _, user := range models.Users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}
