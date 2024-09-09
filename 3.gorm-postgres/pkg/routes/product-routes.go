package routes

import (
	"github.com/gorilla/mux"
	"gorm-postgres/pkg/controllers"
)

var RegisterProductStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/product", controllers.UpdateProduct).Methods("PUT")
}
