package controllers

import (
  "net/http"
  // "strconv"
  "github.com/gin-gonic/gin"
  // "github.com/xiaocuixt/weblo/database"
  // "github.com/xiaocuixt/weblo/lib"
)

func NewSession(c *gin.Context) {
  c.HTML(http.StatusOK, "sessions/new.tmpl", gin.H{
    "title": "user login",
  })
}