package main

import (
  "fmt"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/models"

  "github.com/joho/godotenv"
)

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
  Data.DB.AutoMigrate(&Models.User{})
}
