package repository

import "database/sql"

type Product struct {
	ID       int    `json:"id"`
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func GetProductBySKU(db *sql.DB, sku string) (Product, error) {
	var product Product
	err := db.QueryRow("SELECT id, name, category FROM products WHERE sku = ?", sku).
		Scan(&product.ID, &product.Name, &product.Category)
	product.SKU = sku
	if err != nil {
		return product, err
	}

	return product, nil
}

func GetAllProduct(db *sql.DB) ([]Product, error) {
	var products []Product
	result, err := db.Query("SELECT id, sku, name, category FROM products LIMIT 10")
	if err != nil {
		return products, err
	}

	for result.Next() {
		var product Product
		err := result.Scan(&product.ID, &product.SKU, &product.Name, &product.Category)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}

func CreateProduct(db *sql.DB, product Product) (Product, error) {
	result, err := db.Exec("INSERT INTO products (sku, name, category) VALUES (?, ?, ?)",
		product.SKU, product.Name, product.Category)
	if err != nil {
		return product, err
	}

	id, _ := result.LastInsertId()
	product.ID = int(id)

	return product, nil
}

func UpdateProduct(db *sql.DB, product Product) (Product, error) {
	_, err := db.Exec("UPDATE products SET name = ?, category = ? WHERE sku = ?",
		product.Name, product.Category, product.SKU)
	if err != nil {
		return product, err
	}

	return product, nil
}

func DeleteProduct(db *sql.DB, sku string) error {
	_, err := db.Exec("DELETE FROM products WHERE sku = ?", sku)
	if err != nil {
		return err
	}

	return nil
}
