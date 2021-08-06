// main.go
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/xiaocuixt/weblo/database"
  "github.com/xiaocuixt/weblo/routers"
)

var router *gin.Engine

func main() {
  database.InitDb()
  router := routers.InitRouter()
  router.Run()
}
