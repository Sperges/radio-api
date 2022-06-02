package services

import (
	"errors"
	"radio-api/models"
)

func CreateUser(user models.User) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func ReadUser(id int) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func UpdateUser(id int, newUser models.User) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func DeleteUser(id int) error {
	return errors.New("not implemented")
}
