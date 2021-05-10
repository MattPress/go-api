package main

import (
	"example.com/go-api/controllers"
	"example.com/go-api/models"
	"github.com/gin-gonic/gin"
)

// https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
