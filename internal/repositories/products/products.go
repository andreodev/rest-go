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

func (p Products) GetAll(limit, offset int) ([]productModels.Product, error) {
	rows, err := p.db.Query(`
	SELECT id, name_product, price, description, created_at
	from products
	ORDER BY created_at ASC
	LIMIT $1 OFFSET $2
	`, limit, offset)

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

func (p Products) Count() (int, error) {
	var total int

	if err := p.db.QueryRow(`SELECT COUNT(*) FROM products`).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (p *Products) Create(newProduct productModels.Product) error {

	_, err := p.db.Exec(`
	  INSERT INTO products (id, name_product, price, description, created_at)
	  VALUES ($1, $2, $3, $4, $5)
	`, newProduct.ID, newProduct.NameProduct, newProduct.Price, newProduct.Description, newProduct.CreatedAt,
	)

	return err
}
