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
	err := tx.Commit()
	return err
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
