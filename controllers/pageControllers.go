package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LandingPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetUserPage(name string) gin.HandlerFunc {
	f := func(c *gin.Context) {
		c.HTML(http.StatusOK, name+".html", gin.H{})
	}
	return gin.HandlerFunc(f)
}
