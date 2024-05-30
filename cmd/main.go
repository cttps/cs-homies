package main

import (
	"github.com/cttps/lofi-site/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("frontend/templates/**/*")
	r.Static("frontend/static", "./frontend/static")

	users := []string{"evan"}

	r.GET("/", controllers.LandingPage)

	for _, v := range users {
		r.GET("/"+v, controllers.GetUserPage(v))
	}

	r.Run(":6969")
}
