package Model

import (
  "gorm.io/gorm"
)

type Room struct {
  gorm.Model
  Title string `gorm:"type:varchar(255)"`
  Desc string `gorm:"type:text"`
  X int `gorm:"type:int"`
  Y int `gorm:"type:int"`
  Exits string `gorm:"type:varchar(255)"`
}

func GetUserSessionsInRoom(x int, y int) {

}

func NewRoom(DB *gorm.DB, x int, y int, name string, desc string) (Room, error) {
  r := Room{
    Title: name,
    Desc: desc,
    X: x,
    Y: y,
    Exits: "",
  }

  result := DB.Create(&r)

  if result.Error != nil {
    return r, result.Error
  }

  return r, nil
}
