package Model

import (
  "golang.org/x/crypto/bcrypt"

  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Username string `gorm:"type:varchar(255);unique"`
  Password string `gorm:"type:varchar(255)"`
  Gold int `gorm:"type:int"`
  IsAdmin bool `gorm:"type:bool"`
  IsBuilder bool `gorm:"type:bool"`
}

func NewUser(DB *gorm.DB, username string, password string) (User, error) {
  hashedPassword, _ := HashPassword(password)

  u := User{
    Username: username,
    Password: hashedPassword,
    Gold: 100,
    IsAdmin: false,
    IsBuilder: false,
  }

  result := DB.Create(&u)

  if result.Error != nil {
    return u, result.Error
  }

  return u, nil
}

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func ValidatePassword(password string, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
