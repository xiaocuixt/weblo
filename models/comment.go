package models

import (
   "gorm.io/gorm"
)

type Comment struct {
  gorm.Model

  Content     string `gorm:"type:text"`
  ArticleID   uint
  UserID      uint
  VotesCount  int

  Votes []Vote `gorm:"polymorphic:Votable;polymorphicValue:master"`
}