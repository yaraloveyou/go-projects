package main

import (
	"github.com/gorilla/mux"
	"go-mongodb/db"
	"go-mongodb/routes"
	"log"
	"net/http"
)

func main() {
	err := db.Connect()
	defer db.Disconnect()
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
