package sql

import (
	"eshop/internal/domain"

	"fmt"

	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user *domain.User) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (account,password, type) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Account, user.Password, user.Type)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Auth) GetUser(account, password string) (*domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE account=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, account, password)

	return &user, err
}
