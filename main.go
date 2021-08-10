// main.go
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/routers"
  "github.com/xiaocuixt/weblo/lib"
)

func init() {
  lib.Run()
  database.InitDb()
}
var router *gin.Engine

func main() {
  router := routers.InitRouter()
  router.Run()
}
