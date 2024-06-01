package main

import (
	"os"

	"github.com/cttps/cs-homies/controllers"

	"github.com/cttps/lofi-site/controllers"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin.SetMode(gin.ReleaseMode)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("frontend/templates/**/*")
	r.Static("frontend/static", "./frontend/static")

	r.GET("/", controllers.LandingPage)

	for _, v := range controllers.UserArray {
		r.GET("/"+v, controllers.GetUserPage(v))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)
}
