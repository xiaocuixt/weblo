package controllers

import (
  "net/http"
  // "strconv"
  "github.com/gin-contrib/sessions"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/models"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/lib"
)

func NewSession(c *gin.Context) {
  c.HTML(http.StatusOK, "sessions/new.tmpl", gin.H{
    "title": "user login",
  })
}

func CreateSession(c *gin.Context) {
  var user models.User
  email := c.PostForm("email")
  password := c.PostForm("password")

  database.DB.Where("email = ?", email).First(&user)
  user_password := user.Password

  if lib.CheckPasswordHash(password, user_password) {
    session := sessions.Default(c)
    session.Clear()
    session.Set("userID", user.ID)
    session.Save()
    c.Redirect(http.StatusFound, "/articles")
  }
}

func DeleteSession(c *gin.Context) {
  s := sessions.Default(c)
  s.Clear()
  s.Save()
  c.Redirect(http.StatusFound, "/")
}