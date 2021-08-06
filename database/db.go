package database

import (
   // "fmt"
   "gorm.io/driver/mysql"
   "gorm.io/gorm"
   "github.com/xiaocuixt/weblo/models"
   "github.com/spf13/viper"
   _ "github.com/xiaocuixt/weblo/lib"
)

var DB *gorm.DB

func InitDb() (*gorm.DB, error) {
   dsn := viper.GetString("database.dbuser") +":"+ viper.GetString("database.dbpassword") +"@tcp"+ "(" + viper.GetString("database.dbhost") + ":" + viper.GetString("database.dbport") +")/" + viper.GetString("database.dbname") + "?" + "parseTime=true&loc=Local"
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
