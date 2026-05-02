package products

import productModels "rest-go/internal/models/products"

type ProductRepository interface {
	GetAll() ([]productModels.Product, error)
}
