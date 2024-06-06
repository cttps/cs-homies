package main

import (
	"os"

	"github.com/cttps/cs-homies/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	// Comment if you want to see server debug logs
	// gin.SetMode(gin.ReleaseMode)

	// Uncomment if you want to use a .env file environment variables
	// if err := godotenv.Load(); err != nil {
	// 	panic(err)
	// }

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.LoadHTMLGlob("frontend/templates/**/*")
	r.Static("frontend/static", "./frontend/static")

	r.GET("/", controllers.LandingPage)
	r.GET("/75h", controllers.SevenFive)

	for _, v := range controllers.UserArray {
		r.GET("/"+v, controllers.GetUserPage(v))
	}

	r.POST("/api/verify-user", controllers.VerifyUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run("0.0.0.0:" + port)

}
