package routes

import (
	"radio-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:userid", controllers.ReadUser)
	r.GET("/users", controllers.ReadUsers)
	r.PUT("/users", controllers.UpdateUser)
	r.DELETE("/users/:userid", controllers.DeleteUser)

	r.POST("/:userid/radios", controllers.CreateRadio)
	// r.GET("/:user/radios/:id", controllers.ReadRadio)
	r.GET("/:userid/radios", controllers.ReadRadios)
	r.PUT("/:userid/radios", controllers.UpdateRadio)
	r.DELETE("/:userid/radios/:radioid", controllers.DeleteRadio)

	r.POST("/:userid/playlists", controllers.CreatePlaylist)
	r.POST("/:userid/playlists/:playlistid/radios/:radioid", controllers.AddRadioToPlaylist)
	//r.POST("/users/:user/playlists/:playlistid/radios", controllers.AddRadiosToPlaylist)
	r.GET("/:userid/playlists", controllers.ReadPlaylists)
	r.PUT("/:userid/playlists", controllers.UpdatePlaylist)
	r.DELETE("/:userid/playlists/:playlistid", controllers.DeletePlaylist)
	r.DELETE("/:userid/playlists/:playlistid/radios/:radioid", controllers.RemoveRadioFromPlaylist)
	//r.DELETE("/users/:user/playlists/:playlistid/radios", controllers.RemoveRadiosFromPlaylist)

}
