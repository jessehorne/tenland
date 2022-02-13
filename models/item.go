package Model

import (
  "gorm.io/gorm"
)

type Item struct {
  gorm.Model
  Name string `gorm:"type:varchar(255)"`
  X int `gorm:"type:int"`
  Y int `gorm:"type:int"`
  Held bool `gorm:"type:bool"`
  UserID uint `gorm:"type:int"`
  Description string `gorm:"type:text"`
  Weight float32 `gorm:"type:float;default:0.0"`
}

func ItemNew(DB *gorm.DB, name string, x int, y int, held bool, userID uint) (Item, error) {
  i := Item{
    Name: name,
    X: x,
    Y: y,
    Held: held,
    UserID: userID,
  }

  result := DB.Create(&i)

  if result.Error != nil {
    return i, result.Error
  }

  return i, nil
}

// Transfer ownership of item to user
func ItemHold(DB *gorm.DB, id int, userID uint) bool {
  searchItem := Item{}
  result := DB.First(&searchItem, id)

  if result.RowsAffected == 0 {
    return false
  }

  searchItem.Held = true
  searchItem.UserID = userID

  DB.Save(&searchItem)

  return true;
}

// Drop (or move) item to certain place
func ItemDrop(DB *gorm.DB, id int, x int, y int) bool {
  searchItem := Item{}
  result := DB.First(&searchItem, id)

  if result.RowsAffected == 0 {
    return false
  }

  searchItem.Held = false
  searchItem.X = x
  searchItem.Y = y

  DB.Save(&searchItem)

  return true
}

// Gets all items at X,Y that are not being held
func ItemsGetByRoom(DB *gorm.DB, x int, y int) []Item {
  var items []Item
  DB.Where(&Item{X: x, Y: y, Held: false}).Find(&items)

  return items
}

func ItemsGetByUserID(DB *gorm.DB, userID uint) []Item {
  var items []Item
  DB.Where(&Item{UserID: userID, Held: true}).Find(&items)

  return items
}

func ItemUpdateDescription(DB *gorm.DB, itemID uint, description string) bool {
  searchItem := Item{}
  result := DB.First(&searchItem, itemID)

  if result.RowsAffected == 0 {
    return false
  }

  searchItem.Description = description

  DB.Save(&searchItem)

  return true
}
