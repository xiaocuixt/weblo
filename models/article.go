package models

import (
  "gorm.io/gorm"
  "github.com/go-ozzo/ozzo-validation/v4"
	// "github.com/go-ozzo/ozzo-validation/v4/is"
)

type Article struct {
  gorm.Model
  Title  string
  Content string

  Comments []Comment
}

func (a Article) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.Content, validation.Required.Error("内容不能为空")),
	)
}
