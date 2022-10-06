package controllers

import (
	"api-rest-go/database"
	"api-rest-go/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "Home Page")
	if err != nil {
		return
	}
}

func GetAll(w http.ResponseWriter, _ *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		panic(err.Error())
	}
}

func GetOneById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personality
	database.DB.First(&p, id)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Create(&personality)
	json.NewEncoder(w).Encode(personality)
}

func Delete(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personality
	database.DB.Delete(&p, id)
}
