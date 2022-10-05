package controllers

import (
	"api-rest-go/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "Home Page")
	if err != nil {
		return
	}
}

func GetAll(w http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(w).Encode(models.Personalities)
	if err != nil {
		panic(err.Error())
	}
}

func GetOneById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.Id) == id {
			err := json.NewEncoder(w).Encode(personality)
			if err != nil {
				return
			}
		}
	}
}
