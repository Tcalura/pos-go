package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product // has many
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int // has one
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3307)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// category = Category{Name: "Cozinha"}
	// db.Create(&category)

	// category = Category{Name: "Banho"}
	// db.Create(&category)


	// create product
	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      2999.0,
	// 	CategoryID: 1,
	// })
	// db.Create(&Product{
	// 	Name:       "Panela",
	// 	Price:      99.0,
	// 	CategoryID: 2,
	// })
	// db.Create(&Product{
	// 	Name:       "Toalha",
	// 	Price:      99.0,
	// 	CategoryID: 3,
	// })

	// create serial number
	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1,
	// })

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 2,
	})
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 3,
	})

	var categories []Category
	// para conseguir vizualizar os produtos de cada categoria, precisa usar o Preload
	// e para pegar o serial number, precisa fazer um Preload aninhado 
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
