package controllers

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/models"
  "github.com/gin-contrib/sessions"
)

func ListArticle(c *gin.Context) {
  var articles []models.Article
  database.DB.Find(&articles)

  session := sessions.Default(c)
  userID := session.Get("userID")
  println(userID)

  //c.JSON(http.StatusOK, gin.H{"data": articles})
  c.HTML(http.StatusOK, "articles/index.tmpl", gin.H{
    "articles": articles,
  })
}

func NewArticle(c *gin.Context) {
  c.HTML(http.StatusOK, "articles/new.tmpl", gin.H{
    "title": "new article",
  })
}


func ShowArticle(c *gin.Context) {
  id := c.Param("id")

  var article models.Article
  err := database.DB.First(&article, id).Error
  if err != nil {
    // handle 404
    return
  }

  // c.JSON(http.StatusOK, gin.H{"data": article})
  c.HTML(http.StatusOK, "articles/show.tmpl", gin.H{
    "article": article,
  })
}

func CreateArticle(c *gin.Context) {
  title := c.PostForm("title")
  content := c.PostForm("content")

  article := models.Article{Title: title, Content: content}
  database.DB.Create(&article)

  c.Redirect(http.StatusFound, "/articles")
}

func EditArticle(c *gin.Context) {
  id := c.Param("id")

  var article models.Article
  database.DB.First(&article, id)

  c.HTML(http.StatusOK, "articles/edit.tmpl", gin.H{
    "article": article,
  })
}

func UpdateArticle(c *gin.Context) {
  title := c.PostForm("title")
  content := c.PostForm("content")
  id := c.Param("id")

  var article models.Article
  database.DB.First(&article, id)

  article.Title = title
  article.Content = content
  database.DB.Save(&article)

  c.Redirect(http.StatusFound, "/articles/" + strconv.Itoa(article.ID))
}

func DeleteArticle(c *gin.Context) {
  id := c.Param("id")
  database.DB.Delete(&models.Article{}, id)

  c.JSON(http.StatusOK, gin.H{"data": true})

  // curl -X "DELETE" http://localhost:8080/articles/3
  // {"data":true}
}