package services

import (
	"radio-api/models"
	"radio-api/mysql"
)

func CreateRadio(userId int, radio models.Radio) (int, error) {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("INSERT INTO Radios (RadioName, UserId) VALUES (?, ?)",
		radio.Name,
		userId,
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

func ReadRadio(id int) (models.Radio, error) {
	radio := models.Radio{}
	err := mysql.Database.Get(&radio, "SELECT * FROM Radios WHERE Id=?", id)
	return radio, err
}

func ReadRadiosByUserId(userId int) ([]models.Radio, error) {
	radios := []models.Radio{}
	err := mysql.Database.Select(&radios, "SELECT * FROM Radios WHERE UserId=?", userId)
	return radios, err
}

func ReadRadioIdsByUserId(userId int) ([]int, error) {
	return ReadIdsFromTableByUserId("Radios", userId)
}

func UpdateRadio(radio models.Radio) error {
	if _, err := ReadRadio(radio.Id); err != nil {
		return err
	}
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("UPDATE Radios SET RadioName=? WHERE Id=?",
		radio.Name,
		radio.Id,
	)
	return tx.Commit()
}

func DeleteRadio(id int) error {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("DELETE FROM Radios WHERE Id=?", id)
	return tx.Commit()
}
