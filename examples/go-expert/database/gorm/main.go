package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID    uuid.UUID `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "host=localhost port=5432 user=fullcycle_user password=p@ssw0rd dbname=fullcycle sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	// db.Create(&Product{
	// 	ID:    uuid.New(),
	// 	Name:  "product one",
	// 	Price: 100.0,
	// })

	// create batch
	// products := []Product{
	// 	{ID: uuid.New(), Name: "product one", Price: 100.0},
	// 	{ID: uuid.New(), Name: "product two", Price: 100.0},
	// 	{ID: uuid.New(), Name: "product three", Price: 100.0},
	// }
	// db.Debug().Create(&products)

	// find one by id
	// id, err := uuid.Parse("11c5379c-55ee-480a-8b9a-5d9acc0f2478")
	// if err != nil {
	// 	panic(err)
	// }
	// var product Product
	// db.First(&product, id)
	// fmt.Println(product)

	// find one by name
	// db.First(&product, "name = ?", "product two")
	// fmt.Println(product)

	// find all
	var products []Product
	db.Debug().Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}
}
