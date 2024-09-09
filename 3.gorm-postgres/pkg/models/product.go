package models

import (
	"gorm-postgres/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Name   string  `gorm:"type:varchar(20);not null"`
	Price  float32 `gorm:"type:float;not null"`
	Status int     `gorm:"type:int;not null'"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
}

func (p *Product) CreateProduct() *Product {
	db.Create(&p)
	return p
}

func GetAllProducts() []Product {
	var products []Product
	db.Find(&products)
	return products
}

func GetProductById(id int) *Product {
	var product Product
	db.Where("id = ?", id).First(&product)
	return &product
}

func DeleteProduct(id int) {
	db.Delete(&Product{}, id)
}

func (p *Product) UpdateProduct() *Product {
	db.Save(&p)
	return p
}
