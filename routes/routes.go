package routes

import (
	"api-rest-go/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetAll).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetOneById).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.Create).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.Delete).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.Update).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
