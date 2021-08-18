package models

import (
  // "fmt"
  "errors"
   "gorm.io/gorm"
   "github.com/xiaocuixt/weblo/database"
)

type Comment struct {
  gorm.Model
  Content     string `gorm:"type:text"`
  ArticleID   uint
  UserID      uint
  VotesCount  int

  Votes []Vote `gorm:"polymorphic:Votable;polymorphicValue:master"`
}

func (c Comment) VotedBy(userId uint) bool {
  var votes []Vote

  result := database.DB.Where("votable_id = ? AND votable_type = ? AND user_id = ?", c.ID, "Comment", userId).First(&votes)
  return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
