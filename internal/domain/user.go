package domain

type UserType string

const (
	UserTypeManager UserType = "manager"
	UserTypeUser    UserType = "user"
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	Account  string `json:"account" db:"account"`
	Password string `json:"password" db:"password"`
	Type     string `json:"type" db:"type"` // [‘manager’, ‘user’]
}

func NewUser(account, password string) *User {
	return &User{
		Account:  account,
		Password: password,
		Type:     string(UserTypeUser),
	}
}
