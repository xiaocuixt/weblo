// main.go
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/models"
  "github.com/xiaocuixt/weblo/routers"
  "github.com/xiaocuixt/weblo/lib"
)

func init() {
  lib.Run()
  database.InitDb()
  database.DB.AutoMigrate(&models.Article{})
  database.DB.AutoMigrate(&models.User{})
  database.DB.AutoMigrate(&models.Comment{})
  database.DB.AutoMigrate(&models.Vote{})
}
var router *gin.Engine

func main() {
  router := routers.InitRouter()
  router.Run()
}
