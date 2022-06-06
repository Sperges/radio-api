package services

import (
	"fmt"
	"radio-api/mysql"
)

func ReadIdsFromTableByUserId(table string, userId int) ([]int, error) {
	pStatement := fmt.Sprintf("SELECT Id FROM %s WHERE UserId=?", table)
	ids := []int{}
	err := mysql.Database.Select(&ids, pStatement, userId)
	return ids, err
}
