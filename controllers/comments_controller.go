package controllers

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/models"
  "github.com/gin-contrib/sessions"
)

func CreateComment(c *gin.Context) {
  articleId := c.PostForm("article_id")
  // user_id := c.PostForm("user_id")
  _articleId, _ := strconv.ParseUint(articleId, 10, 64)
  content := c.PostForm("content")

  session := sessions.Default(c)
  sessionUserID := session.Get("userID")
  userID := sessionUserID.(uint)

  comment := models.Comment{ArticleID: uint(_articleId), UserID: userID, Content: content}
  //comment := models.Comment{Content: content}
  database.DB.Create(&comment)

  c.Redirect(http.StatusFound, "/articles/" + articleId)
}