package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/lokeshajja/rest_api_gorm_postgres/models"
  "github.com/lokeshajja/rest_api_gorm_postgres/controllers"
  // "models"
)


func main() {
  r := gin.Default()

  models.ConnectDatabase()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.GET("/books", controllers.FindBooks )
  r.POST("/books", controllers.CreateBook)
  r.GET("/books/:id", controllers.FindBook) // new
  r.PATCH("/books/:id", controllers.UpdateBook) // new
  r.DELETE("/books/:id", controllers.DeleteBook) // new





  r.Run()
}