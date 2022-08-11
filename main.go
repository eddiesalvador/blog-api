package main

import (
	_ "fmt"
	"net/http"

	"blogAPI/controllers"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

func main() {

	r := setupRouter()
	// _ = r.Run(":8080")

	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*") // loads the templates and the html files
	r.Static("/css", "./css/")       // needed to load css files for the site
	r.StaticFile("favicon.ico", "media/img/favicon.ico") // (go) relative path == (HTML) href link

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.GET("/templates/pages", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "DevBlog | Home",
		})
	})

	r.GET("/templates/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about", gin.H{
			"title": "DevBlog | About",
		})
	})

	userRepo := controllers.New()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	return r
}
