package main

import (
	"os"

	"github.com/cttps/lofi-site/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil { // loading environment var from file to be used by os.Getenv
		panic(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("frontend/templates/**/*")
	r.Static("frontend/static", "./frontend/static")

	users := []string{"bobby", "evan"}

	r.GET("/", controllers.LandingPage)

	for _, v := range users {
		r.GET("/"+v, controllers.GetUserPage(v))
	}

	r.Run(os.Getenv("LISTEN_ADDY"))
}
