package routes

import (
	"radio-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.ReadUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.POST("/radio", controllers.CreateRadio)
	r.GET("/radio/:id", controllers.ReadRadio)
	r.GET("/:user/radios/:id", controllers.ReadRadioByIDFromUser)
	r.PUT("/radio/:id", controllers.UpdateRadio)
	r.DELETE("/radio/:id", controllers.DeleteRadio)

	r.POST("/playlist", controllers.CreatePlaylist)
	r.GET("/playlist/:id", controllers.ReadPlaylist)
	r.GET("/:user/playlists/:id", controllers.ReadPlaylistByIDFromUser)
	r.PUT("/playlist/:id", controllers.UpdatePlaylist)
	r.DELETE("/playlist/:id", controllers.DeletePlaylist)

}
