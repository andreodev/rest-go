package usecases

import (
	"log/slog"
	productModels "rest-go/internal/models/products"
	productRepo "rest-go/internal/repositories/products"
)

type ProductsUseCase struct {
	repo productRepo.ProductRepository
}

func NewProductsUseCase(repo productRepo.ProductRepository) *ProductsUseCase {
	return &ProductsUseCase{repo: repo}
}

func (u ProductsUseCase) GetAll() ([]productModels.Product, error) {
	products, err := u.repo.GetAll()
	if err != nil {
		slog.Error("failed to get products", "err", err)
		return nil, err
	}

	return products, nil
}
