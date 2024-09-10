package routes

import (
	"github.com/gorilla/mux"
	"go-mongodb/controllers"
)

var RegisterRoutes = func(r *mux.Router) {
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
}
