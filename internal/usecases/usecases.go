package usecases

import (
	"rest-go/internal/repositories"
)

type UseCases struct {
	Users    *UsersUseCase
	Products *ProductsUseCase
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		Users:    NewUsersUseCase(repos.User),
		Products: NewProductsUseCase(repos.Products),
	}
}
