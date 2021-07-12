package database

import (
   // "fmt"
   "gorm.io/driver/mysql"
   "gorm.io/gorm"
   "github.com/xiaocuixt/weblo/models"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "weblo"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var DB *gorm.DB

func InitDb() (*gorm.DB, error) {
   dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
   if err != nil {
     panic("failed to connect database")
   }

   DB = db
   db.AutoMigrate(&models.Article{})
   db.AutoMigrate(&models.User{})
   db.AutoMigrate(&models.Comment{})
   db.AutoMigrate(&models.Vote{})

   return db, err
}