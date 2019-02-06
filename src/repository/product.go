package repository

type Product struct {
	ID       int    `json:"id"`
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func GetProductBySKU(sku string) Product {
	return Product{
		SKU:      "abc123",
		Name:     "Cardinal Casual",
		Category: "Men - Casual - Trousers",
	}
}
