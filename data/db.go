package Data

import (
  "os"
  "strings"

  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

var dsn string = "<USERNAME>:<PASSWORD>@tcp(127.0.0.1:3306)/<DBNAME>?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() {
  dsn = strings.Replace(dsn, "<USERNAME>", os.Getenv("MYSQL_USER"), 1)
  dsn = strings.Replace(dsn, "<PASSWORD>", os.Getenv("MYSQL_PASS"), 1)
  dsn = strings.Replace(dsn, "<DBNAME>", os.Getenv("MYSQL_DB"), 1)

  DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
