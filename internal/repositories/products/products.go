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

func (p Products) FindByID(id string) (productModels.Product, error) {
	var product productModels.Product

	err := p.db.QueryRow(`
		SELECT id, name_product, price, description, created_at
		FROM products
		WHERE id = $1
	`, id).Scan(
		&product.ID,
		&product.NameProduct,
		&product.Price,
		&product.Description,
		&product.CreatedAt,
	)

	return product, err
}

func (p *Products) DeleteByID(id string) error {
	result, err := p.db.Exec(`
	  DELETE FROM products
	  WHERE id = $1
	`, id)

	if err != nil {
		return err
	}

	row, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if row == 0 {
		return sql.ErrNoRows
	}

	return nil
}
