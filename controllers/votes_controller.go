package controllers

import (
  "net/http"
  "strconv"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/models"
  "github.com/gin-contrib/sessions"
)

func CreateVote(c *gin.Context) {
  var comment models.Comment

  commentId := c.PostForm("comment_id")
  _commentId, _ := strconv.Atoi(commentId)
  err := database.DB.First(&comment, _commentId).Error
  if err != nil {
    return
  }
  voteType := c.PostForm("vote_type")
  _voteType, _ := strconv.Atoi(voteType)

  session := sessions.Default(c)
  sessionUserID := session.Get("userID")
  userID := sessionUserID.(uint)

  vote := models.Vote{
    VotableID: _commentId,
    VotableType: "Comment",
    UserID: userID,
    VoteType: _voteType,
  }
  result := database.DB.Create(&vote)
  fmt.Println(result.Error)

  if result.Error == nil {
    if _voteType > 0 {
      database.DB.Model(&comment).Update("VotesCount", comment.VotesCount + 1)
    } else {
      database.DB.Model(&comment).Update("VotesCount", comment.VotesCount - 1)
    }
  }
  c.Redirect(http.StatusFound, "/articles/" + strconv.Itoa(int(comment.ArticleID)))
}
