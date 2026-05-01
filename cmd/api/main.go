package main

import (
	"rest-go/internal/handlers"
	"rest-go/internal/repositories"
	"rest-go/internal/usecases"
)

func main() {
	repos := repositories.New()

	useCases := usecases.New(repos)

	h := handlers.New(useCases)

	h.Listen(8080)
}
