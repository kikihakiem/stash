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

func GetAll() []Product {
	return []Product{
		Product{
			SKU:      "abc123",
			Name:     "Cardinal Casual",
			Category: "Men - Casual - Trousers",
		},
		Product{
			SKU:      "abc124",
			Name:     "Cardinal Formal",
			Category: "Men - Casual - Trousers",
		},
	}
}
