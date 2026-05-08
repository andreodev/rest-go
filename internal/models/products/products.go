package products

import (
	"errors"
	"math"
	"strings"

	"github.com/google/uuid"
)

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

func (r CreateProductRequest) ValidateProduct() error {

	if strings.TrimSpace(r.NameProduct) == "" {
		return errors.New("name product is required")
	}

	if math.IsNaN(r.Price) || math.IsInf(r.Price, 0) {
		return errors.New("invalid price")
	}

	if r.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if strings.TrimSpace(r.Description) == "" {
		return errors.New("description product is required")
	}

	return nil
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

type ProductDeleteResponse struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}

type ProductUpdateRequest struct {
	NameProduct string  `json:"name_product"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ProductUpdateResponse struct {
	ID          uuid.UUID `json:"id"`
	NameProduct string    `json:"name_product"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
}
