package services

import (
	"errors"
	"radio-api/models"
)

func CreateRadio(Radio models.Radio) (models.Radio, error) {
	return models.Radio{}, errors.New("not implemented")
}

func ReadRadio(id int) (models.Radio, error) {
	return models.Radio{}, errors.New("not implemented")
}

func ReadRadioByIDFromUser(user models.User, id int) (models.Radio, error) {
	return models.Radio{}, errors.New("not implemented")
}

func ReadRadioByNameFromUser(user models.User, id int) (models.Radio, error) {
	return models.Radio{}, errors.New("not implemented")
}

func UpdateRadio(id int, newRadio models.Radio) (models.Radio, error) {
	return models.Radio{}, errors.New("not implemented")
}

func DeleteRadio(id int) error {
	return errors.New("not implemented")
}
