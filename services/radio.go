package services

import (
	"fmt"
	"radio-api/models"
	"radio-api/mysql"
)

func CreateRadio(userId int, radio models.Radio) (int, error) {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("INSERT INTO Radios (RadioName, StreamUrl, UserId) VALUES (?, ?, ?)", radio.Name, radio.StreamUrl, userId)
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return GetNewestId()
}

func ReadRadiosByUserId(userId int) ([]models.Radio, error) {
	radios := []models.Radio{}
	err := mysql.Database.Select(&radios, "SELECT * FROM Radios WHERE UserId=?", userId)
	for i := range radios {
		playlistIds, err := ReadPlaylistIdsByRadioId(radios[i].Id)
		if err != nil {
			return radios, err
		}
		radios[i].PlaylistIds = playlistIds
	}
	return radios, err
}

func ReadRadioIdsByPlaylistId(playlistId int) ([]int, error) {
	ids := []int{}
	err := mysql.Database.Select(&ids, "SELECT RadioId FROM PlaylistRadios WHERE PlaylistId=?", playlistId)
	return ids, err
}

func ReadRadioIdsByUserId(userId int) ([]int, error) {
	ids := []int{}
	err := mysql.Database.Select(&ids, "SELECT Id FROM Radios WHERE UserId=?", userId)
	return ids, err
}

func UpdateRadio(userId int, radio models.Radio) error {
	fmt.Println(radio)
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	result := tx.MustExec("UPDATE Radios SET RadioName=?, StreamUrl=? WHERE Id=? AND UserId=?",
		radio.Name,
		radio.StreamUrl,
		radio.Id,
		userId,
	)
	if err := CheckUpdateSuccess(result); err != nil {
		return err
	}
	return tx.Commit()
}

func DeleteRadio(userId int, id int) error {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("DELETE FROM Radios WHERE Id=? AND UserId=?", id, userId)
	return tx.Commit()
}
