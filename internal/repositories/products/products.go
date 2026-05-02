package products

import (
	"database/sql"
	productModels "rest-go/internal/models/products"
)

type Products struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Products {
	return &Products{db: db}
}

func (p Products) GetAll() ([]productModels.Product, error) {
	rows, err := p.db.Query(`
	SELECT id, name_product, price, description, created_at
	from products
	ORDER BY created_at ASC
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]productModels.Product, 0)

	for rows.Next() {
		var product productModels.Product

		if err := rows.Scan(
			&product.ID,
			&product.NameProduct,
			&product.Price,
			&product.Description,
			&product.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
