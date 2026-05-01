package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"rest-go/internal/usecases"
)

type Handlers struct {
	useCases usecases.UseCases
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers{useCases: *useCases}
}

func (h Handlers) Listen(port int) error {
	h.registerUserEndpoints()

	slog.Info("server is listening", "port", port)

	return http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		nil,
	)
}
