package services

import (
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

func UpdateUser(user models.User) error {
	if _, err := repo.ReadUser(user.Id); err != nil {
		return err
	}
	return repo.UpdateUser(user)
}

func DeleteUser(id int) error {
	return repo.DeleteUser(id)
}
