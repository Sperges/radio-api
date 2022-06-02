package services

import (
	"errors"
	"radio-api/models"
)

func CreatePlaylist(Playlist models.Playlist) (models.Playlist, error) {
	return models.Playlist{}, errors.New("not implemented")
}

func ReadPlaylist(id int) (models.Playlist, error) {
	return models.Playlist{}, errors.New("not implemented")
}

func ReadPlaylistByIDFromUser(user models.User, id int) (models.Playlist, error) {
	return models.Playlist{}, errors.New("not implemented")
}

func ReadPlaylistByNameFromUser(user models.User, id int) (models.Playlist, error) {
	return models.Playlist{}, errors.New("not implemented")
}

func UpdatePlaylist(id int, newPlaylist models.Playlist) (models.Playlist, error) {
	return models.Playlist{}, errors.New("not implemented")
}

func DeletePlaylist(id int) error {
	return errors.New("not implemented")
}
