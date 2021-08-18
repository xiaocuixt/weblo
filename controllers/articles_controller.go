package controllers

import (
  "fmt"
  "net/http"
  "strconv"
  "html/template"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/models"
  "github.com/xiaocuixt/weblo/utils"
  "reflect"
  "math"
)

func ListArticle(c *gin.Context) {
  var count int64
  var articles []models.Article
  database.DB.Model(&models.Article{}).Count(&count)
  pages, _ := strconv.Atoi(c.Query("per_page"))
  if pages <= 0 {
    pages = 10
  }
  currentPage, _ := strconv.Atoi(c.Query("page"))
  totalPages := int(math.Ceil(float64(count) / float64(pages)))
  fmt.Println(count)
  fmt.Println(totalPages)
  database.DB.Scopes(utils.Paginate(c.Query("page"), c.Query("per_page"), c)).Find(&articles)

  //c.JSON(http.StatusOK, gin.H{"data": articles})
  c.HTML(http.StatusOK, "articles/index.tmpl", gin.H{
    "articles": articles,
    "totalPages": totalPages,
    "currentPage": currentPage,
  })
}

func NewArticle(c *gin.Context) {
  c.HTML(http.StatusOK, "articles/new.tmpl", gin.H{})
}

func ShowArticle(c *gin.Context) {
  id := c.Param("id")
  user, _ := c.Get("currentUser")
  var article models.Article
  var comments []models.Comment
  // var voteableIds []string

  err := database.DB.First(&article, "id = ?", id).Error
  if err != nil {
    // handle 404
    return
  }

  err = database.DB.Where("article_id = ?", id).Find(&comments).Error
  if err != nil {
    // handle 404
    return
  }
  // voteIDs := database.DB.Model(&models.Vote{}).Where("user_id = ?", user.Email).Distinct().Pluck("voteable_id", &voteableIds)
  // fmt.Println(voteIDs)
  // myVote := database.DB.Where("user_id = ?", user.ID).Find(&votes)
  // hasVotedArticle := 

  // c.JSON(http.StatusOK, gin.H{"data": article})
  c.HTML(http.StatusOK, "articles/show.tmpl", gin.H{
    "article": article,
    "content": template.HTML(article.Content),
    "currentUser": user,
    "comments": comments,
  })
}

func CreateArticle(c *gin.Context) {
  title := c.PostForm("title")
  content := c.PostForm("content")

  article := models.Article{Title: title, Content: content}
  err := article.Validate()
  fmt.Println(reflect.TypeOf(err))
  fmt.Println(err)
  var errClass string
  if err != nil {
    errClass = "error"
    c.HTML(http.StatusOK, "articles/new.tmpl", gin.H{
			"article": article,
      "message": err,
      "errClass": errClass,
		})
		return
  }
  
  database.DB.Create(&article)

  c.Redirect(http.StatusFound, "/dashboard/articles")
}

func EditArticle(c *gin.Context) {
  id := c.Param("id")

  var article models.Article
  database.DB.First(&article, "id = ?", id)

  c.HTML(http.StatusOK, "articles/edit.tmpl", gin.H{
    "article": article,
  })
}

func UpdateArticle(c *gin.Context) {
  title := c.PostForm("title")
  content := c.PostForm("content")
  id := c.Param("id")
  fmt.Println(id)
  var article models.Article
  database.DB.First(&article, "id = ?", id)

  article.Title = title
  article.Content = content

  fmt.Println(title)

  validate_err := article.Validate()
  fmt.Println(reflect.TypeOf(validate_err))
  fmt.Println(validate_err)
  var errClass string
  if validate_err != nil {
    errClass = "error"
    c.HTML(http.StatusOK, "articles/new.tmpl", gin.H{
			"article": article,
      "message": validate_err,
      "errClass": errClass,
		})
		return
  }

  database.DB.Save(&article)

  // c.Redirect(http.StatusFound, "/dashboard/articles/" + strconv.Itoa(article.ID))
  c.Redirect(http.StatusFound, "/dashboard/articles")
}

func DeleteArticle(c *gin.Context) {
  id := c.Param("id")
  database.DB.Delete(&models.Article{}, id)

  c.JSON(http.StatusOK, gin.H{"data": true})

  // curl -X "DELETE" http://localhost:8080/articles/3
  // {"data":true}
}
