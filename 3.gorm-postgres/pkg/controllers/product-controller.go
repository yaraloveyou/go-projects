package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm-postgres/pkg/models"
	"gorm-postgres/pkg/utils"
	"net/http"
	"strconv"
)

var NewProduct models.Product

func GetProducts(w http.ResponseWriter, r *http.Request) {
	newProducts := models.GetAllProducts()
	res, _ := json.Marshal(newProducts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	productId := param["id"]
	id, err := strconv.Atoi(productId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	productDetails := models.GetProductById(id)
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	productId := param["id"]
	id, err := strconv.Atoi(productId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	models.DeleteProduct(id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}
	err := utils.ParseBody(r, product)
	b := product.CreateProduct()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}
	err := utils.ParseBody(r, product)
	b := product.UpdateProduct()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
