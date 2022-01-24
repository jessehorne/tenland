package Game

import (
  "golang.org/x/crypto/bcrypt"
)

type Player struct {
  Username string
  Password string

  Gold int
}

func NewPlayer(username string, password string) Player {
  hashedPassword, _ := HashPassword(password)

  p := Player{
    Username: username,
    Password: hashedPassword,
    Gold: 100,
  }

  return p
}

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func ValidatePassword(password string, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
