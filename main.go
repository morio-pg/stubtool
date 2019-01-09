package main

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := InitializeRouter().Setup()
	router.Static("/css", "./dist/css")
	router.Static("/fonts", "./dist/fonts")
	router.Static("/img", "./dist/img")
	router.Static("/js", "./dist/js")
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
		} else {
			c.File("./dist/index.html")
		}
	})
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://stubtool.herokuapp.com", "http://localhost:8080"}
	router.Use(cors.New(config))
	router.Run()
}
