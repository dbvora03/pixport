package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() error {
   var err error
   DB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/photos"), &gorm.Config{})
   if err != nil {
       return err
   }
   return nil
}