package usecases

import (
	"rest-go/internal/repositories"
	authUseCase "rest-go/internal/usecases/auth"
)

type UseCases struct {
	Users    *UsersUseCase
	Products *ProductsUseCase
	Auth     *authUseCase.AuthUseCase
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		Users:    NewUsersUseCase(repos.User),
		Products: NewProductsUseCase(repos.Products),
		Auth:     authUseCase.NewAuthUseCase(repos.Auth),
	}
}
