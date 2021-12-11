package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// get db
func GetConnction(host string, port int, user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=5s", user, password, host, port, dbname)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "db connect err")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "connect close")
	}

	return db, err
}
