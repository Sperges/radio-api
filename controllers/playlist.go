package controllers

import (
	"net/http"
	"radio-api/models"
	"radio-api/services"

	"github.com/gin-gonic/gin"
)

func CreatePlaylist(c *gin.Context) {
	var playlist models.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if result, err := services.CreatePlaylist(playlist); err != nil {
		_ = c.AbortWithError(http.StatusNotImplemented, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func ReadPlaylist(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if result, err := services.ReadPlaylist(id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func ReadPlaylists(c *gin.Context) {

}

func ReadPlaylistByIDFromUser(c *gin.Context) {

}

func ReadPlaylistByNameFromUser(c *gin.Context) {

}

func UpdatePlaylist(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	var playlist models.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if result, err := services.UpdatePlaylist(id, playlist); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func DeletePlaylist(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if err := services.DeletePlaylist(id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, "")
	}
}
