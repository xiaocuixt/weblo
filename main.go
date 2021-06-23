// main.go
package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
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