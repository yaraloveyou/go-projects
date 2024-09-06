package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Price        float32       `json:"price"`
	Manufacturer *Manufacturer `json:"manufacturer"`
}

type Manufacturer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var products []Product

func main() {
	var PORT string = ":8080"
	r := mux.NewRouter()

	products = append(products, Product{
		ID:    1,
		Name:  "Cheaps",
		Price: 1.2,
		Manufacturer: &Manufacturer{
			ID:   1,
			Name: "Lays",
		},
	})

	products = append(products, Product{
		ID:    2,
		Name:  "Crackers",
		Price: 0.9,
		Manufacturer: &Manufacturer{
			ID:   2,
			Name: "Fishka",
		},
	})

	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	fmt.Printf("Starting server at port: %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range products {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode("Product not found")
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = len(products) + 1
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
		return
	}
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	for i, item := range products {
		if item.ID == id {
			products = append(products[:i], products[i+1:]...)
			product.ID = id
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	json.NewEncoder(w).Encode("Product not found")
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, item := range products {
		if item.ID == id {
			products = append(products[:id], products[id+1:]...)
		}
	}
}
