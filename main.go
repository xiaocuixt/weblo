// main.go
package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/models"
)

var router *gin.Engine

func main() {
  Db := database.InitDb()
  Db.AutoMigrate(&models.Article{})

  router := gin.Default()
  // loads all the template files located in the templates folder
  router.LoadHTMLGlob("templates/**/*")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
      "title": "Weblo",
    })
  })

  router.GET("/articles", func(c *gin.Context) {
    c.HTML(http.StatusOK, "articles/index.tmpl", gin.H{
      "title": "articles list",
    })
  })

  router.GET("/articles/new", func(c *gin.Context) {
    c.HTML(http.StatusOK, "articles/new.tmpl", gin.H{
      "title": "new article",
    })
  })

  router.Run()
}