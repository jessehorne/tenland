package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

func DropCommandHandler(cmd []string, session *Game.Session) {
  // Get users current position
  x := session.User.X
  y := session.User.Y

  // Get room at current user position
  searchRoom := Model.Room{}
  result := Data.DB.Where(Model.Room{
    X: x,
    Y: y,
  }).First(&searchRoom)

  // Handle if room doesn't exist in database
  if result.RowsAffected == 0 {
    session.Conn.Write([]byte("You can't drop an item into the abyss.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  // Make sure argument is provided
  if len(cmd) < 2 {
    session.Conn.Write([]byte("You have to specify what you want to drop.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  // Get list of items in room
  allItems := Model.ItemsGetByUserID(Data.DB, session.User.ID)

  for _,v := range allItems {
    if v.Name == cmd[1] && v.Held {
      // Item exists, now let user pick it up
      v.Held = false
      v.X = x
      v.Y = y
      Data.DB.Save(&v)

      session.Conn.Write([]byte("You've dropped " + v.Name + ".\n"))
      session.Conn.Write([]byte(Data.Cursor))

      return
    }
  }

  session.Conn.Write([]byte("You can't drop something that doesn't exist. Or can you?\n"))
  session.Conn.Write([]byte(Data.Cursor))
}

func NewDropCommand() CommandType {
  hc := NewCommand("drop", "'drop <name>' - Drop an item from your inventory.")
  hc.Handler = DropCommandHandler
  AllCommandsBig["drop"] =
  "Usage: 'drop <name>'\n" +
  "Drop an item from your inventory.\n"

  return hc
}
