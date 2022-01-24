package main

import (
  "fmt"

  "github.com/jessehorne/tenland/data"

  "gorm.io/gorm"

  "github.com/joho/godotenv"
)

type User struct {
  gorm.Model
  Username string `gorm:"type:varchar(255);unique"`
  Password string `gorm:"type:varchar(255)"`
}

func main() {
  // Environment Variables
  // Load environment variables
  err := godotenv.Load()

  if err != nil {
    fmt.Println("[ERROR] Error loading .env file...")
  }

  // Initialize Database
  Data.InitDB()

  // Migrate
  Data.DB.AutoMigrate(&User{})
}
