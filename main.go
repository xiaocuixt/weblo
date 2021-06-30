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

  router := gin.Default()

  router.Static("/assets", "./assets")
  // router.Static("/vendor", "./assets/vendor")
  // r.Static("/img", "./static/img")
  // r.Static("/scss", "./static/scss")
  // r.Static("/vendor", "./static/vendor")
  // r.Static("/js", "./static/js")
  // r.StaticFile("/favicon.ico", "./img/favicon.ico")

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