package services

import (
	"radio-api/models"
	"radio-api/mysql"
)

func CreatePlaylist(userId int, playlist models.Playlist) (int, error) {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("INSERT INTO Playlists (PlaylistName, UserId) VALUES (?, ?)",
		playlist.Name,
		userId,
	)
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return GetNewestId()
}

func ReadPlaylist(id int) (models.Playlist, error) {
	var playlist models.Playlist
	err := mysql.Database.Get(&playlist, "SELECT * FROM Playlists WHERE Id=?", id)
	return playlist, err
}

func ReadPlaylistsByUserId(userId int) ([]models.Playlist, error) {
	var playlists []models.Playlist
	err := mysql.Database.Select(&playlists, "SELECT * FROM Playlists WHERE UserId=?", userId)
	for i := range playlists {
		radioIds, err := ReadRadioIdsByPlaylistId(playlists[i].Id)
		if err != nil {
			return playlists, err
		}
		playlists[i].RadioIds = radioIds
	}
	return playlists, err
}

func ReadPlaylistIdsByRadioId(radioId int) ([]int, error) {
	ids := []int{}
	err := mysql.Database.Select(&ids, "SELECT PlaylistId FROM PlaylistRadios WHERE RadioId=?", radioId)
	return ids, err
}

func ReadPlaylistIdsByUserId(userId int) ([]int, error) {
	ids := []int{}
	err := mysql.Database.Select(&ids, "SELECT Id FROM Playlists WHERE UserId=?", userId)
	return ids, err
}

func UpdatePlaylist(userId int, playlist models.Playlist) error {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	result := tx.MustExec("UPDATE Playlists SET PlaylistName=? WHERE Id=? AND UserId=?",
		playlist.Name,
		playlist.Id,
		userId,
	)
	if err := CheckUpdateSuccess(result); err != nil {
		return err
	}
	return tx.Commit()
}

func DeletePlaylist(userId int, id int) error {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("DELETE FROM Playlists WHERE Id=? AND UserId=?", id, userId)
	return tx.Commit()
}

func AddRadioToPlaylist(userId int, playlistId int, radioId int) error {
	if _, err := ReadUser(userId); err != nil {
		return err
	}
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("INSERT INTO PlaylistRadios (PlaylistId, RadioId, UserId) VALUES (?, ?, ?)", playlistId, radioId, userId)
	return tx.Commit()
}

func RemoveRadioFromPlaylist(userId int, playlistId int, radioId int) error {
	tx := mysql.Database.MustBegin()
	defer tx.Rollback()
	tx.MustExec("DELETE FROM PlaylistRadios WHERE PlaylistId=? AND RadioId=? AND UserId=?", playlistId, radioId, userId)
	return tx.Commit()
}

// func AddRadiosToPlaylist(playlistId int, radioIds []int) error {
// 	return errors.New("not implemented")
// }

// func RemoveRadiosFromPlaylist(playlistId int, radioIds []int) error {
// 	return errors.New("not implemented")
// }
