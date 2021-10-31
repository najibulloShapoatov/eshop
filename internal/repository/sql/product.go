package sql

import (
	"eshop/internal/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	db *sqlx.DB
}

func NewProduct(db *sqlx.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Create(product *domain.Product) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (name,description, quantity) values ($1, $2, $3) RETURNING id", productsTable)
	row := p.db.QueryRow(query, product.Name, product.Description, product.Quantity)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (p *Product) GetByID(id int64) (*domain.Product, error) {
	query := fmt.Sprintf("select * from %s where id =$1", productsTable)
	var product domain.Product
	err := p.db.Get(&product, query, id)
	return &product, err
}

func (p *Product) GetList() ([]*domain.Product, error) {
	query := fmt.Sprintf("select * from %s", productsTable)
	var products []*domain.Product
	err := p.db.Select(&products, query)
	return products, err
}

func (p *Product) Update(id int64, form *domain.Product) error {
	query := fmt.Sprintf("update %s set name=$1, description=$2, is_active=$3, updated_at=$4, quantity=$5 where id=$6", productsTable)
	_, err := p.db.Exec(query, form.Name, form.Description, form.IsActive, form.UpdatedAt, form.Quantity, id)
	return err
}

func (p *Product) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id =$1", productsTable)
	_, err := p.db.Exec(query, id)
	return err
}
