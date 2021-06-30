package controllers

import (
  "net/http"
  // "strconv"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/models"
  "github.com/xiaocuixt/weblo/lib"
)

func NewUser(c *gin.Context) {
  c.HTML(http.StatusOK, "users/new.tmpl", gin.H{
    "title": "user register",
  })
}

func CreateUser(c *gin.Context) {
  email := c.PostForm("email")
  password, err := lib.HashPassword(c.PostForm("password"))
  if err != nil {
    return
  }
  // TODO add email and password validator later

  user := models.User{Email: email, Password: password}
  database.DB.Create(&user)

  c.Redirect(http.StatusFound, "/users/login")
}