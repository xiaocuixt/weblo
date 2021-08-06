package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/xiaocuixt/weblo/models"
	"github.com/xiaocuixt/weblo/database"
	"github.com/gin-contrib/sessions"
)

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
func AuthRequired() gin.HandlerFunc {
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
