package products

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	NameProduct string    `json:"name_product"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at"`
}
