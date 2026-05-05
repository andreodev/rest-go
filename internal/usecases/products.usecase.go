package usecases

import (
	"log/slog"
	productModels "rest-go/internal/models/products"
	productRepo "rest-go/internal/repositories/products"
	"time"

	"github.com/google/uuid"
)

type ProductsUseCase struct {
	repo productRepo.ProductRepository
}

func NewProductsUseCase(repo productRepo.ProductRepository) *ProductsUseCase {
	return &ProductsUseCase{repo: repo}
}

func (p ProductsUseCase) GetAll(page, limit int) (productModels.GetAllProductsResponse, error) {
	offset := (page - 1) * limit

	products, err := p.repo.GetAll(limit, offset)
	if err != nil {
		slog.Error("failed to get products", "err", err)
		return productModels.GetAllProductsResponse{}, err
	}

	total, err := p.repo.Count()
	if err != nil {
		slog.Error("failed to count products", "err", err)
		return productModels.GetAllProductsResponse{}, err
	}

	return productModels.GetAllProductsResponse{
		Products: products,
		Page:     page,
		Limit:    limit,
		Total:    total,
	}, nil
}

func (p ProductsUseCase) Create(newProduct productModels.CreateProductRequest) (productModels.CreateProductResponse, error) {
	repoReq := productModels.Product{
		ID:          uuid.New(),
		NameProduct: newProduct.NameProduct,
		Price:       newProduct.Price,
		Description: newProduct.Description,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	if err := p.repo.Create(repoReq); err != nil {
		slog.Error("failed to create product", "name_product", newProduct.NameProduct, "err", err)
		return productModels.CreateProductResponse{}, err
	}

	return productModels.CreateProductResponse{ID: repoReq.ID}, nil
}

func (p ProductsUseCase) FindByID(id string) (productModels.Product, error) {
	product, err := p.repo.FindByID(id)

	if err != nil {
		slog.Error("Failed to get product by id", "id", id, "err", err)

		return productModels.Product{}, err
	}

	return product, nil
}

func (p ProductsUseCase) DeleteByID(id string) error {
	if err := p.repo.DeleteByID(id); err != nil {
		slog.Error("FAILED TO DELETE PRODUCT", "id", id, "err", err)
		return err
	}

	return nil
}
