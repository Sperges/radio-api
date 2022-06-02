package controllers

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.POST("/user", CreateUser)
	r.GET("/user/:id", ReadUser)
	r.PUT("/user/:id", UpdateUser)
	r.DELETE("/user/:id", DeleteUser)

	r.POST("/radio", CreateRadio)
	r.GET("/radio/:id", ReadRadio)
	r.GET("/:user/radios/:id", ReadRadioByIDFromUser)
	r.GET("/:user/radios/:name", ReadRadioByNameFromUser)
	r.PUT("/radio/:id", UpdateRadio)
	r.DELETE("/radio/:id", DeleteRadio)

	r.POST("/playlist", CreatePlaylist)
	r.GET("/playlist/:id", ReadPlaylist)
	r.GET("/:user/playlists/:id", ReadPlaylistByIDFromUser)
	r.GET("/:user/playlists/:name", ReadPlaylistByNameFromUser)
	r.PUT("/playlist/:id", UpdatePlaylist)
	r.DELETE("/playlist/:id", DeletePlaylist)

}
