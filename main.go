// main.go
package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/controllers"

  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
)

var router *gin.Engine

func main() {
  database.InitDb()

  router := gin.Default()
  store := cookie.NewStore([]byte("secret"))
  router.Use(sessions.Sessions("weblo", store))

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

  // user auth
  router.GET("/users/signup", controllers.NewUser)
  router.POST("/users", controllers.CreateUser)
  router.GET("/users/login", controllers.NewSession)
  router.POST("/users/login", controllers.CreateSession)

  router.Run()
}