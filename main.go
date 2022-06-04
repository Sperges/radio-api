package main

import (
	"os"
	"radio-api/mysql"
	"radio-api/routes"

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

	routes.SetupRoutes(r)

	r.Run(host + ":" + port)
}
