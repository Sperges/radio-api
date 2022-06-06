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

func ReadPlaylistIdsByUserId(userId int) ([]int, error) {
	return ReadIdsFromTableByUserId("Playlists", userId)
}

func UpdatePlaylist(id int, newPlaylist models.Playlist) (models.Playlist, error) {
	return models.Playlist{}, errors.New("not implemented")
}

func DeletePlaylist(id int) error {
	return errors.New("not implemented")
}
