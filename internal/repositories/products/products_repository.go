package products

import productModels "rest-go/internal/models/products"

type ProductRepository interface {
	GetAll(limit, offset int) ([]productModels.Product, error)
	Count() (int, error)
	Create(productModels.Product) error
	FindByID(id string) (productModels.Product, error)
}
