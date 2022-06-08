package controllers

import (
	"radio-api/models"
	"radio-api/services"

	"github.com/gin-gonic/gin"
)

func CreatePlaylist(c *gin.Context) {
	userId, paramErr := GetPathParamAsInt(c, "userid")
	if paramErr != nil {
		return
	}
	var playlist models.Playlist
	if bindErr := BindStruct(c, &playlist); bindErr != nil {
		return
	}
	result, resultErr := services.CreatePlaylist(userId, playlist)
	StructResponse(c, result, resultErr)
}

func ReadPlaylists(c *gin.Context) {
	userId, paramErr := GetPathParamAsInt(c, "userid")
	if paramErr != nil {
		return
	}
	result, resultErr := services.ReadPlaylistsByUserId(userId)
	StructResponse(c, result, resultErr)
}

func UpdatePlaylist(c *gin.Context) {
	userId, paramErr := GetPathParamAsInt(c, "userid")
	if paramErr != nil {
		return
	}
	var playlist models.Playlist
	if bindErr := BindStruct(c, &playlist); bindErr != nil {
		return
	}
	OkResponse(c, services.UpdatePlaylist(userId, playlist))
}

func DeletePlaylist(c *gin.Context) {
	userId, paramErr := GetPathParamAsInt(c, "userid")
	if paramErr != nil {
		return
	}
	id, paramErr := GetPathParamAsInt(c, "id")
	if paramErr != nil {
		return
	}
	OkResponse(c, services.DeletePlaylist(userId, id))
}

func AddRadioToPlaylist(c *gin.Context) {
	userId, paramErr := GetPathParamAsInt(c, "userid")
	if paramErr != nil {
		return
	}
	playlistId, paramErr := GetPathParamAsInt(c, "playlistid")
	if paramErr != nil {
		return
	}
	radioId, paramErr := GetPathParamAsInt(c, "radioid")
	if paramErr != nil {
		return
	}
	OkResponse(c, services.AddRadioToPlaylist(userId, playlistId, radioId))
}

func RemoveRadioFromPlaylist(c *gin.Context) {
	userId, paramErr := GetPathParamAsInt(c, "userid")
	if paramErr != nil {
		return
	}
	playlistId, paramErr := GetPathParamAsInt(c, "playlistid")
	if paramErr != nil {
		return
	}
	radioId, paramErr := GetPathParamAsInt(c, "radioid")
	if paramErr != nil {
		return
	}
	OkResponse(c, services.RemoveRadioFromPlaylist(userId, playlistId, radioId))
}

// func AddRadiosToPlaylist(c *gin.Context) {
// 	_, paramErr := GetPathParamAsInt(c, "userid")
// 	if paramErr != nil {
// 		return
// 	}
// 	playlistId, paramErr := GetPathParamAsInt(c, "playlistid")
// 	if paramErr != nil {
// 		return
// 	}
// 	radioIds := []int{}
// 	OkResponse(c, services.RemoveRadiosFromPlaylist(playlistId, radioIds))
// }

// func RemoveRadiosFromPlaylist(c *gin.Context) {
// 	_, paramErr := GetPathParamAsInt(c, "userid")
// 	if paramErr != nil {
// 		return
// 	}
// 	playlistId, paramErr := GetPathParamAsInt(c, "playlistid")
// 	if paramErr != nil {
// 		return
// 	}
// 	radioIds := []int{}
// 	OkResponse(c, services.RemoveRadiosFromPlaylist(playlistId, radioIds))
// }
