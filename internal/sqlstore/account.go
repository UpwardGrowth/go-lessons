package sqlstore

import "database/sql"

type Userinfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Account struct {
	DB *sql.DB
}

func New(conn *sql.DB) *Account {
	return &Account{conn}
}

func (ac Account) QueryUserinfo(id string) (Userinfo, error) {
	var user Userinfo
	err := ac.DB.QueryRow("SELECT username, email FROM users where id=?", id).Scan(&user.Name, &user.Email)
	if err != nil {
		return Userinfo{}, err
	}
	return user, nil
}
