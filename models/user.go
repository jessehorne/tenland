package Models

import (
  "golang.org/x/crypto/bcrypt"

  "gorm.io/gorm"

  "github.com/jessehorne/tenland/data"
)

type User struct {
  gorm.Model
  Username string `gorm:"type:varchar(255);unique"`
  Password string `gorm:"type:varchar(255)"`
  Gold int `gorm:"type:int"`
}

func NewUser(username string, password string) User {
  hashedPassword, _ := HashPassword(password)

  u := User{
    Username: username,
    Password: hashedPassword,
    Gold: 100,
  }

  Data.DB.Create(&u)

  return u
}

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func ValidatePassword(password string, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
