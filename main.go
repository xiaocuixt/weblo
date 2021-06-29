// main.go
package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/controllers"
)

var router *gin.Engine

func main() {
  database.InitDb()
  println("start program ....... ")

  router := gin.Default()
  // loads all the template files located in the templates folder
  router.LoadHTMLGlob("templates/**/*")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
      "title": "Weblo",
    })
  })

  router.GET("/articles", controllers.ListArticle)
  router.GET("/articles/:id", controllers.ShowArticle)
  router.GET("/articles/new", controllers.NewArticle)
  router.POST("/articles", controllers.CreateArticle)
  router.GET("/articles/:id/edit", controllers.EditArticle)
  router.POST("/articles/:id", controllers.UpdateArticle)
  router.DELETE("/articles/:id", controllers.DeleteArticle)

  router.Run()
}