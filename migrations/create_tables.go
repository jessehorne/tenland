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

  userTableError := Data.DB.AutoMigrate(&Model.User{})

  if userTableError != nil {
    fmt.Println("Error migrating users table.")
  }

  roomTableError := Data.DB.AutoMigrate(&Model.Room{})

  if roomTableError != nil {
    fmt.Println("Error migrating rooms table.")
  }

  configTableError := Data.DB.AutoMigrate(&Model.Config{})

  if configTableError != nil {
    fmt.Println("Error migrating configs table.")
  }

  itemTableError := Data.DB.AutoMigrate(&Model.Item{})

  if itemTableError != nil {
    fmt.Println("Error migrating items table.")
  }
}
