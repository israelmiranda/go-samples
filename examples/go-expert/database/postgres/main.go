package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func findAll(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM public.products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func findOne(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM public.products WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func insert(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO public.products(id, name, price) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func update(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE public.products SET name = $1, price = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func delete(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM public.products WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=fullcycle_user password=p@ssw0rd dbname=fullcycle sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productOne := NewProduct("product one", 1000.05)

	// insert
	err = insert(db, productOne)
	if err != nil {
		panic(err)
	}

	// update
	productOne.Price = 100.0
	err = update(db, productOne)
	if err != nil {
		panic(err)
	}

	// find one
	p, err := findOne(db, productOne.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v - $%.2f\n", p.Name, p.Price)

	// find all
	products, err := findAll(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("%v - $%.2f\n", p.Name, p.Price)
	}

	// delete
	err = delete(db, productOne.ID)
	if err != nil {
		panic(err)
	}
}
