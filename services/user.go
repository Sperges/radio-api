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
	return GetNewestId()
}

func ReadUser(id int) (models.User, error) {
	var user models.User
	err := mysql.Database.Get(&user, "SELECT * FROM Users WHERE Id=?", id)
	if err != nil {
		return user, err
	}
	radioIds, playlistIds, err := populateUserIds(user.Id)
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
	for i := range users {
		radioIds, playlistIds, err := populateUserIds(users[i].Id)
		if err != nil {
			return users, err
		}
		users[i].RadioIds = radioIds
		users[i].PlaylistIds = playlistIds
	}
	return users, err
}

func UpdateUser(user models.User) error {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	result := tx.MustExec("UPDATE Users SET Username=? WHERE Id=?",
		user.Name,
		user.Id,
	)
	if err := CheckUpdateSuccess(result); err != nil {
		return err
	}
	return tx.Commit()
}

func DeleteUser(id int) error {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("DELETE FROM Users WHERE Id=?", id)
	return tx.Commit()
}

// Populate radio and playlist ids.
// radioId, playlistId, error.
func populateUserIds(userId int) (radioIds []int, playlistIds []int, err error) {
	radioIds, err = ReadRadioIdsByUserId(userId)
	if err != nil {
		return []int{}, []int{}, err
	}
	playlistIds, err = ReadPlaylistIdsByUserId(userId)
	if err != nil {
		return []int{}, []int{}, err
	}
	return radioIds, playlistIds, nil
}
