// main.go
package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/models"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/controllers"
  // "github.com/xiaocuixt/weblo/lib"
  // "fmt"
  // "github.com/spf13/viper"

  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
)

var router *gin.Engine

func main() {
  database.InitDb()
  // fmt.Println(viper.GetString("name"))
  router := gin.Default()
  store := cookie.NewStore([]byte("secret"))
  router.Use(sessions.Sessions("weblo", store))
  router.Use(GetCurrentUser())

  router.Static("/assets", "./assets")

  // loads all the template files located in the templates folder
  router.LoadHTMLGlob("templates/**/*")
  authorized := router.Group("/")
  authorized.Use(authRequired())
  {
    authorized.POST("/comments", controllers.CreateComment)
    authorized.POST("/votes", controllers.CreateVote)
  }

  dashboard := router.Group("/dashboard")
  dashboard.Use(authRequired())
  {
    dashboard.GET("/", func(c *gin.Context) {
      user, _ := c.Get("currentUser")
      c.HTML(http.StatusOK, "dashboard/index.tmpl", gin.H{
        "title": "Dashboard",
        "currentUser": user,
      })
    })
    dashboard.GET("/articles", controllers.ListArticle)
    dashboard.GET("/articles/new", controllers.NewArticle)
    dashboard.GET("/articles/:id/edit", controllers.EditArticle)
    dashboard.POST("/articles", controllers.CreateArticle)
    dashboard.POST("/articles/:id", controllers.UpdateArticle)
    dashboard.DELETE("/articles/:id", controllers.DeleteArticle)
  }

  router.GET("/", func(c *gin.Context) {
    var articles []models.Article
    database.DB.Find(&articles)
    user, _ := c.Get("currentUser")
    c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
      "title": "Weblo",
      "currentUser": user,
      "articles": articles,
    })
  })

  // router.Use(checkCurrentUser())  全局middleware
  router.GET("/articles/:id", controllers.ShowArticle)

  // user auth
  router.GET("/users/signup", controllers.NewUser)
  router.POST("/users", controllers.CreateUser)
  router.GET("/users/login", controllers.NewSession)
  router.POST("/users/login", controllers.CreateSession)
  router.GET("/users/logout", controllers.DeleteSession)

  router.Run()
}


func GetCurrentUser() gin.HandlerFunc {
  return func(c *gin.Context) {
    session := sessions.Default(c)
    var user models.User
    if userID := session.Get("userID"); userID != nil {
      err := database.DB.First(&user, userID).Error
      if err == nil {
        c.Set("currentUser", user)
      }
    }
    c.Next()
  }
}

// currentUser := c.MustGet("currentUser")
func authRequired() gin.HandlerFunc {
  return func(c *gin.Context) {
    var user models.User
    session := sessions.Default(c)
    userID := session.Get("userID")
    err := database.DB.First(&user, userID).Error

    if err != nil {
      c.Redirect(http.StatusFound, "/users/login")
    } else {
      c.Next()
    }
  }
}
