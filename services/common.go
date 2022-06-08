package services

import (
	"database/sql"
	"radio-api/mysql"
)

func GetNewestId() (int, error) {
	var id int
	if err := mysql.Database.Get(&id, "SELECT LAST_INSERT_ID()"); err != nil {
		return 0, err
	}
	return id, nil
}

func CheckUpdateSuccess(result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}
	return nil
}
