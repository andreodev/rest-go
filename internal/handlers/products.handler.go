package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"rest-go/internal/models"
	productModels "rest-go/internal/models/products"
	"strconv"

	"github.com/google/uuid"
)

func (h Handlers) registerProductEndpoints() {
	http.HandleFunc("GET /products", h.requireAuth(h.getAllProducts))
	http.HandleFunc("GET /products/", h.requireAuth(h.getAllProducts))
	http.HandleFunc("POST /products", h.requireAuth(h.createProduct))
	http.HandleFunc("PUT /products/{id}", h.requireAuth(h.updateProductByID))
	http.HandleFunc("GET /products/{id}", h.requireAuth(h.getProductByID))
	http.HandleFunc("DELETE /products/{id}", h.requireAuth(h.deleteProductByID))
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

	if err := req.ValidateProduct(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
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

func (p Handlers) getProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := getProductIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	product, err := p.useCases.Products.FindByID(id.String())

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "product not found"})

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func getProductIDFromRequest(r *http.Request) (uuid.UUID, error) {
	id := r.PathValue("id")

	if id != "" {
		return uuid.Parse(id)
	}

	var req productModels.ProductByIDResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return uuid.Nil, errors.New("invalid request body")
	}

	if req.ID == uuid.Nil {
		return uuid.Nil, errors.New("id is required")
	}

	return req.ID, nil
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

func (h Handlers) deleteProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := getProductIDFromRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})

		return
	}

	if err := h.useCases.Products.DeleteByID(id.String()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "products not found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productModels.ProductDeleteResponse{
		Message: "product deleted successfully",
		ID:      id,
	})
}

func (h Handlers) updateProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := getProductIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	var req productModels.ProductUpdateRequest

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "invalid request body"})
		return
	}

	fmt.Println("Request body:", string(bodyBytes))

	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "invalid request body"})
		return
	}

	if err := h.useCases.Products.UpdateByID(id.String(), req); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "product not found"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productModels.ProductUpdateResponse{ID: id, NameProduct: req.NameProduct, Price: req.Price, Description: req.Description})
}
