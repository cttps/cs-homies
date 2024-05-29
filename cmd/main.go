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
	r.LoadHTMLGlob("frontend/**/*")
	r.Static("frontend/static", "./frontend/static")

	r.GET("/", controllers.LandingPage)

	r.Run(os.Getenv("LISTEN_ADDY"))
}
