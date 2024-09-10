package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go-postgres/handlers"
	"go-postgres/repository"
	"go-postgres/services"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=gorm password=admin sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookRepo := repository.NewBookRepository(db)
	bookServ := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookServ)

	router := mux.NewRouter()
	router.HandleFunc("/book", bookHandler.Create).Methods("POST")
	router.HandleFunc("/book", bookHandler.GetAll).Methods("GET")
	router.HandleFunc("/book", bookHandler.Update).Methods("PUT")
	router.HandleFunc("/book/{id}", bookHandler.Delete).Methods("DELETE")
	router.HandleFunc("/book/{id}", bookHandler.Get).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
