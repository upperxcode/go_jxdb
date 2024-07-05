package models

import (
	"database/sql"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

func scanProduct(rows *sql.Rows) (Product, error) {
	var product Product
	err := rows.Scan(&product.ID, &product.Name, &product.Price)
	return product, err
}

func productValues(product Product) []interface{} {
	return []interface{}{product.Name, product.Price}
}

func productIDValue(product Product) interface{} {
	return product.ID
}
