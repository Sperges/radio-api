package services

import (
	"errors"
	"radio-api/models"
	"radio-api/repo"
)

func CreateUser(user models.User) error {
	return repo.CreateUser(user)
}

func ReadUser(id int) (models.User, error) {
	return repo.ReadUser(id)
}

func ReadUsers() ([]models.User, error) {
	return repo.ReadUsers()
}

func UpdateUser(id int, newUser models.User) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func DeleteUser(id int) error {
	return errors.New("not implemented")
}
