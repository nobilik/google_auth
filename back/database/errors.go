package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func IsDuplicateEntryError(err error) bool {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return false
	}

	return sqlErr.Number == 1062
}

func IsNotFoundError(err error) bool {
	return err == sql.ErrNoRows
}
