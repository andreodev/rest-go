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
	h.registerProductEndpoints()
	h.registerDocsEndpoints()

	slog.Info("server is listening", "port", port)

	return http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		withCors(http.DefaultServeMux),
	)
}

func withCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
