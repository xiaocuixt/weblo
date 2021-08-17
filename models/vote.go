package models

import (
   "gorm.io/gorm"
)

type Vote struct {
  gorm.Model
  UserID      uint   `gorm:"uniqueIndex:idx_user_votable"`
  VoteType   int      //enum, 支持(1)or反对(0)
  VotableID   int     `gorm:"uniqueIndex:idx_user_votable"`
  VotableType string  `gorm:"uniqueIndex:idx_user_votable"`
}
