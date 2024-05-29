package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo: switch around so that landing page = user page, and user page = sign up/log in page

func LandingPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
