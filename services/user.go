package services

import (
	"radio-api/models"
	"radio-api/mysql"
)

func CreateUser(user models.User) (int, error) {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("INSERT INTO Users (Username) VALUES (?)",
		user.Name,
	)
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	var id int
	if err := mysql.Database.Get(&id, "SELECT LAST_INSERT_ID()"); err != nil {
		return 0, err
	}
	return id, nil
}

func ReadUser(id int) (models.User, error) {
	user := models.User{}
	err := mysql.Database.Get(&user, "SELECT * FROM Users WHERE Id=?", id)
	if err != nil {
		return user, err
	}
	radioIds, err := ReadRadioIdsByUserId(id)
	if err != nil {
		return user, err
	}
	playlistIds, err := ReadPlaylistIdsByUserId(id)
	if err != nil {
		return user, err
	}
	user.RadioIds = radioIds
	user.PlaylistIds = playlistIds
	return user, err
}

func ReadUsers() ([]models.User, error) {
	users := []models.User{}
	err := mysql.Database.Select(&users, "SELECT * FROM Users ORDER BY Id ASC")
	return users, err
}

func UpdateUser(user models.User) error {
	if _, err := ReadUser(user.Id); err != nil {
		return err
	}
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("UPDATE Users SET Username=? WHERE Id=?",
		user.Name,
		user.Id,
	)
	return tx.Commit()
}

func DeleteUser(id int) error {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("DELETE FROM Users WHERE Id=?", id)
	return tx.Commit()
}
