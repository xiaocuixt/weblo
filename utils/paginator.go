// 参考: https://github.com/unknwon/paginater/blob/master/paginater.go
//  https://gowalker.org/github.com/astaxie/beego/utils/paginationhttps://github.com/astaxie/beego/blob/master/utils/pagination/paginator.go
package utils

import (
	"strconv"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func Paginate(page string, perPage string, c *gin.Context) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    page, _ := strconv.Atoi(page)
    if page == 0 {
      page = 1
    }

    pageSize, _ := strconv.Atoi(perPage)
    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}
