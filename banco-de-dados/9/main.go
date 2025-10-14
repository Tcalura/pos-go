package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3307)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// chamado lock pessimista
	// lock otimista é o gorm.Model que ja tem o campo Version
	// e toda vez que atualizar o registro, ele incrementa o valor do Version
	// se duas transações tentarem atualizar o mesmo registro ao mesmo tempo
	// a segunda transação vai falhar, pois o valor do Version não vai bater
	// com o valor do Version do banco de dados

	// nesse exemplo, vamos usar o lock pessimista
	// que bloqueia o registro para outras transações
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletronicos testes"
	tx.Debug().Save(&c)
	tx.Commit()
}