package Model

import (
  "golang.org/x/crypto/bcrypt"

  "gorm.io/gorm"
)

/*
  Player

  Stats
  =====
  ST = Strength

  Other
  =====
  Max Carry Weight = ST + 100
*/

type User struct {
  gorm.Model
  Username string `gorm:"type:varchar(255);unique"`
  Password string `gorm:"type:varchar(255)"`
  GoldBank int `gorm:"type:int"`
  GoldHand int `gorm:"type:int"`
  IsAdmin bool `gorm:"type:bool"`
  IsBuilder bool `gorm:"type:bool"`
  X int `gorm:"type:int"`
  Y int `gorm:"type:int"`
  ST int `gorm:"type:int;default:1"`
}

func (self User) GetMaxCarryWeight() int {
  return self.ST + 100
}

func NewUser(DB *gorm.DB, username string, password string) (User, error) {
  hashedPassword, _ := HashPassword(password)

  u := User{
    Username: username,
    Password: hashedPassword,
    GoldBank: 0,
    GoldHand: 0,
    IsAdmin: true,
    IsBuilder: true,
    ST: 1,
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
