package main

import (
	"fmt"
	// "net/http"

	// "blogAPI/controllers"
	"blogAPI/database"

	// "github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

func main() {

	// r := setupRouter()
	// _ = r.Run(":8080")
	db := database.InitDb()	// Initializes connection to database

	tablelist, err := db.Migrator().GetTables()
	
	if err != nil {
		fmt.Println("error getting tables")
	}
	fmt.Print("List of table names: ")
	fmt.Println(tablelist)


	// r.Run()
}

// func setupRouter() *gin.Engine {
// 	r := gin.Default()

// 	// r.LoadHTMLGlob("templates/**/*") // loads the templates and the html files
// 	// r.Static("/css", "./css/")       // needed to load css files for the site
// 	// r.StaticFile("favicon.ico", "media/img/favicon.ico") // (go) relative path == (HTML) href link

// 	r.GET("ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, "pong")
// 	})

// 	// r.GET("/", func(c *gin.Context) {
// 	// 	c.HTML(http.StatusOK, "index", gin.H{
// 	// 		"title": "DevBlog | Home",
// 	// 	})
// 	// })

// 	// r.GET("/about", func(c *gin.Context) {
// 	// 	c.HTML(http.StatusOK, "about", gin.H{
// 	// 		"title": "DevBlog | About",
// 	// 	})
// 	// })

// 	userRepo := controllers.New()
// 	r.GET("/users", userRepo.GetUsers)
// 	r.GET("/users/:id", userRepo.GetUser)
// 	r.POST("/users", userRepo.CreateUser)
// 	r.PUT("/users/:id", userRepo.UpdateUser)
// 	r.DELETE("/users/:id", userRepo.DeleteUser)

// 	return r
// }
