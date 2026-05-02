package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest-go/internal/models"
	productModels "rest-go/internal/models/products"
	"strconv"
)

func (h Handlers) registerProductEndpoints() {
	http.HandleFunc("GET /products", h.getAllProducts)
	http.HandleFunc("POST /products", h.createProduct)
}

func (h Handlers) getAllProducts(w http.ResponseWriter, r *http.Request) {
	page, limit, err := getPaginationFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	resp, err := h.useCases.Products.GetAll(page, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h Handlers) createProduct(w http.ResponseWriter, r *http.Request) {
	var req productModels.CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "invalid request body"})
		return
	}

	resp, err := h.useCases.Products.Create(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func getPaginationFromRequest(r *http.Request) (int, int, error) {
	page, err := getPositiveQueryInt(r, "page", 1)
	if err != nil {
		return 0, 0, err
	}

	limit, err := getPositiveQueryInt(r, "limit", 10)
	if err != nil {
		return 0, 0, err
	}

	return page, limit, nil
}

func getPositiveQueryInt(r *http.Request, key string, defaultValue int) (int, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return defaultValue, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil || parsed < 1 {
		return 0, errors.New(key + " must be a positive integer")
	}

	return parsed, nil
}
