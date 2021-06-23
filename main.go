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
  InitDb()
  router := gin.Default()
  // loads all the template files located in the templates folder
  router.LoadHTMLGlob("templates/*")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Weblo",
    })
  })

  router.Run()
}