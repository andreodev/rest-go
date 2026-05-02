package handlers

import (
	"encoding/json"
	"net/http"
	"rest-go/internal/models"
)

func (h Handlers) registerProductEndpoints() {
	http.HandleFunc("GET /products", h.getAllProducts)
}

func (h Handlers) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.useCases.Products.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
