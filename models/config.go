package Model

import (
  "gorm.io/gorm"
)

type Config struct {
  gorm.Model
  Key string `gorm:"type:varchar(255)"`
  Value string `gorm:"type:varchar(255)"`
}

func NewConfig(DB *gorm.DB, key string, value string) (Config, error) {
  c := Config{
    Key: key,
    Value: value,
  }

  result := DB.Create(&c)

  if result.Error != nil {
    return c, result.Error
  }

  return c, nil
}
