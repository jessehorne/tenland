package Model

import (
  "gorm.io/gorm"

  "github.com/jessehorne/tenland/data"
)

type Room struct {
  gorm.Model
  Name string `gorm:"type:varchar(255)"`
  Desc string `gorm:"type:text"`
  X int `gorm:"type:int"`
  Y int `gorm:"type:int"`
}

func NewRoom(x int, y int, name string, desc string) (Room, error) {
  r := Room{
    Name: name,
    Desc: desc,
    X: x,
    Y: y,
  }

  result := Data.DB.Create(&r)

  if result.Error != nil {
    return r, result.Error
  }

  return r, nil
}
