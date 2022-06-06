package routes

import (
	"radio-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.ReadUser)
	r.GET("/users", controllers.ReadUsers)
	r.PUT("/users", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.POST("/:user/radios", controllers.CreateRadio)
	// r.GET("/:user/radios/:id", controllers.ReadRadio)
	r.GET("/:user/radios", controllers.ReadRadios)
	r.PUT("/:user/radios", controllers.UpdateRadio)
	r.DELETE("/:user/radios/:id", controllers.DeleteRadio)

	r.POST("/:user/playlists", controllers.CreatePlaylist)
	r.GET("/:user/playlists/:id", controllers.ReadPlaylist)
	r.GET("/:user/playlists", controllers.ReadPlaylists)
	r.PUT("/:user/playlists", controllers.UpdatePlaylist)
	r.DELETE("/:user/playlists/:id", controllers.DeletePlaylist)

}
