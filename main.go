package main

import (
	"os"
	"radio-api/controllers"
	"radio-api/mysql"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	mysql.Init()
	defer mysql.Close()

	r := gin.Default()
	r.SetTrustedProxies([]string{host})

	controllers.SetupRoutes(r)

	r.Run(host + ":" + port)
}
