package models

import (
   "gorm.io/gorm"
)

type Comment struct {
  gorm.Model

  Content     string `gorm:"type:text"`
  ArticleID   uint
  UserID      uint
}