package models

import (
   "gorm.io/gorm"
)

type Vote struct {
  gorm.Model
  UserID      uint
  VoteType   int //enum, 支持(1)or反对(0)
  VotableID   int
  VotableType string
}
