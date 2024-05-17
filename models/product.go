package models

type Product struct {
	ProductID   int64   `json:"product_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SellerID    int64   `json:"seller_id"`
}

type Products struct {
	Products []Product `json:"products"`
}
