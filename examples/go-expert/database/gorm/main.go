package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// type Product struct {
// 	ID    uuid.UUID `gorm:"primaryKey"`
// 	Name  string
// 	Price float64
// }

// type Product struct {
// 	ID    uuid.UUID `gorm:"primaryKey"`
// 	Name  string
// 	Price float64
// 	gorm.Model
// }

// type Category struct {
// 	ID   uuid.UUID `gorm:"primaryKey"`
// 	Name string
// }

// type Product struct {
// 	ID         uuid.UUID `gorm:"primaryKey"`
// 	Name       string
// 	Price      float64
// 	CategoryID uuid.UUID
// 	Category   Category
// 	gorm.Model
// }

// type Category struct {
// 	ID       uuid.UUID `gorm:"primaryKey"`
// 	Name     string
// 	Products []Product
// }

// type Product struct {
// 	ID           uuid.UUID `gorm:"primaryKey"`
// 	Name         string
// 	Price        float64
// 	CategoryID   uuid.UUID
// 	Category     Category
// 	SerialNumber SerialNumber
// }

type Category struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories"`
}

type Product struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories"`
}

// type SerialNumber struct {
// 	ID        uuid.UUID `gorm:"primaryKey"`
// 	Number    string
// 	ProductID uuid.UUID
// }

// func deleteAll(db *gorm.DB) {
// 	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Product{})
// }

// func createOne(db *gorm.DB) {
// 	db.Create(&Product{
// 		ID:    uuid.New(),
// 		Name:  "product one",
// 		Price: 100,
// 	})
// }

// func createOneWithCategory(db *gorm.DB) {
// 	category := Category{ID: uuid.New(), Name: "category one"}
// 	db.Create(&category)

// 	db.Create(&Product{
// 		ID:         uuid.New(),
// 		Name:       "product one",
// 		Price:      100,
// 		CategoryID: category.ID,
// 	})
// }

// func createBatch(db *gorm.DB) {
// 	products := []Product{
// 		{ID: uuid.New(), Name: "product two", Price: 200},
// 		{ID: uuid.New(), Name: "product three", Price: 300},
// 		{ID: uuid.New(), Name: "product four", Price: 400},
// 		{ID: uuid.New(), Name: "product five", Price: 500},
// 	}
// 	db.Create(&products)
// }

// func createBatchWithCategory(db *gorm.DB) {
// 	category := Category{ID: uuid.New(), Name: "category two"}
// 	db.Create(&category)

// 	products := []Product{
// 		{ID: uuid.New(), Name: "product two", Price: 200, CategoryID: category.ID},
// 		{ID: uuid.New(), Name: "product three", Price: 300, CategoryID: category.ID},
// 		{ID: uuid.New(), Name: "product four", Price: 400, CategoryID: category.ID},
// 		{ID: uuid.New(), Name: "product five", Price: 500, CategoryID: category.ID},
// 	}
// 	db.Create(&products)
// }

// func createOneWithCategoryAndSerialNumber(db *gorm.DB) {
// 	category := Category{ID: uuid.New(), Name: "category one"}
// 	db.Create(&category)

// 	product := Product{
// 		ID:         uuid.New(),
// 		Name:       "product one",
// 		Price:      100,
// 		CategoryID: category.ID,
// 	}
// 	db.Create(&product)

// 	serialNumber := SerialNumber{
// 		ID:        uuid.New(),
// 		Number:    "123456",
// 		ProductID: product.ID,
// 	}
// 	db.Create(&serialNumber)
// }

func main() {
	dsn := "host=localhost port=5432 user=fullcycle_user password=p@ssw0rd dbname=fullcycle sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&Product{})
	// db.AutoMigrate(&Product{}, &Category{})
	// db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	db.AutoMigrate(&Product{}, &Category{})

	// delete all
	// deleteAll(db)

	// create one
	// createOne(db)

	// create batch
	// createBatch(db)

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
	// var products []Product
	// db.Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// find all
	// var products []Product
	// db.Limit(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price >= ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// like
	// var products []Product
	// db.Where("name LIKE ?", "%ne%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update
	// var product Product
	// db.Where("name = ?", "product one").Find(&product)
	// product.Name = "final product"
	// db.Save(&product)

	// delete
	// var p Product
	// db.Where("name = ?", "final product").Find(&p)
	// fmt.Println(p)
	// db.Delete(&p)

	// create one with category
	// createOneWithCategory(db)

	// create batch with category
	// createBatchWithCategory(db)

	// find all with category
	// var products []Product
	// db.Preload("Category").Find(&products) // lazy
	// // db.Joins("Category").Find(&products) // eager
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name)
	// }

	// create one with category and serial number
	// createOneWithCategoryAndSerialNumber(db)

	// find all with category and serial number
	// var products []Product
	// db.Joins("Category").Joins("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	// find all categories with all products
	// var category Category
	// db.Preload("Products").Find(&category)

	// for _, p := range category.Products {
	// 	fmt.Println(p)
	// }
	// var categories []Category
	// err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	// err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	fmt.Println("Category:", category.Name)
	// 	for _, product := range category.Products {
	// 		fmt.Println(
	// 			" - Product:", product.Name,
	// 			"-", product.SerialNumber.Number,
	// 		)
	// 	}
	// }

	// // create category
	// categoryOne := Category{ID: uuid.New(), Name: "category one"}
	// db.Create(&categoryOne)
	// categoryTwo := Category{ID: uuid.New(), Name: "category two"}
	// db.Create(&categoryTwo)

	// // create product
	// db.Create(&Product{
	// 	ID:         uuid.New(),
	// 	Name:       "product one",
	// 	Price:      99.99,
	// 	Categories: []Category{categoryOne, categoryTwo},
	// })

	// // find all many to many
	// var categories []Category
	// db.Preload("Products").Find(&categories)
	// for _, category := range categories {
	// 	fmt.Println("Category:", category.Name)
	// 	for _, product := range category.Products {
	// 		fmt.Println(
	// 			" - Product:", product.Name,
	// 		)
	// 	}
	// }

	// pessimistic locking
	tx := db.Begin()
	var category Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, "name = ?", "category one").Error
	if err != nil {
		panic(err)
	}
	category.Name = "final category"
	tx.Debug().Save(&category)
	tx.Commit()
}
