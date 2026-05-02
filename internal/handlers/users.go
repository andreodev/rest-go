package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"rest-go/internal/models"
	userModels "rest-go/internal/models/users"

	"github.com/google/uuid"
)

func (h Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /users", h.getAllUsers)
	http.HandleFunc("POST /users", h.createUser)
	http.HandleFunc("DELETE /users/{id}", h.deleteUserById)
	http.HandleFunc("GET /users/{id}", h.getUserById)
}

func (h Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.useCases.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h Handlers) createUser(w http.ResponseWriter, r *http.Request) {

	var req userModels.UserCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "Invalid request body"})
		return
	}

	id, err := h.useCases.AddUser(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userModels.UserCreateResponse{NewUserID: id})
}

func (h Handlers) deleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	if err := h.useCases.DeleteUserById(id.String()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "user not found"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userModels.UserDeleteResponse{
		Message: "user deleted successfully",
		ID:      id,
	})
}

func (h Handlers) getUserById(w http.ResponseWriter, r *http.Request) {
	id, err := getUserIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	user, err := h.useCases.GetUserById(id.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "user not found"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func getUserIDFromRequest(r *http.Request) (uuid.UUID, error) {
	id := r.PathValue("id")
	if id != "" {
		return uuid.Parse(id)
	}

	var req userModels.UserDeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return uuid.Nil, errors.New("invalid request body")
	}

	if req.Id == uuid.Nil {
		return uuid.Nil, errors.New("id is required")
	}

	return req.Id, nil
}
