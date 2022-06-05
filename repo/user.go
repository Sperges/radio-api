package repo

import (
	"radio-api/models"
	"radio-api/mysql"
)

func CreateUser(user models.User) error {
	tx := mysql.Database.MustBegin()
	tx.MustExec("INSERT INTO Users (Username) VALUES (?)",
		user.Name,
	)
	return tx.Commit()
}

func ReadUser(id int) (models.User, error) {
	user := models.User{}
	err := mysql.Database.Get(&user, "SELECT * FROM Users WHERE Id=?", id)
	return user, err
}

func ReadUsers() ([]models.User, error) {
	users := []models.User{}
	err := mysql.Database.Select(&users, "SELECT * FROM Users ORDER BY Id ASC")
	return users, err
}

func UpdateUser(user models.User) error {
	tx := mysql.Database.MustBegin()
	tx.MustExec("UPDATE Users SET Username=? WHERE Id=?",
		user.Name,
		user.Id,
	)
	return tx.Commit()
}

func DeleteUser(id int) error {
	tx := mysql.Database.MustBegin()
	tx.MustExec("DELETE FROM Users WHERE Id=?", id)
	return tx.Commit()
}
