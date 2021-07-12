package models

import (
   "gorm.io/gorm"
)

type Article struct {
  gorm.Model
  ID    int
  Title  string
  Content string

  Comments []Comment
}