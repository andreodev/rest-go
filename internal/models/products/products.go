package products

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`
	NameProduct string    `json:"name_product"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at"`
}

type CreateProductRequest struct {
	NameProduct string  `json:"name_product"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type GetAllProductsResponse struct {
	Products []Product `json:"products"`
	Page     int       `json:"page"`
	Limit    int       `json:"limit"`
	Total    int       `json:"total"`
}

type CreateProductResponse struct {
	ID          uuid.UUID `json:"id"`
	NameProduct string    `json:"name_product"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at"`
}

type ProductByIDResponse struct {
	ID          uuid.UUID `json:"id"`
	NameProduct string    `json:"name_product"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at"`
}
