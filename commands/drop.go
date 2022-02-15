package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/arg"
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
    Arg.WriteFull(session.Conn, "You can't drop an item into the abyss.\n")
    return
  }

  // Make sure argument is provided
  if len(cmd) < 2 {
    Arg.WriteFull(session.Conn, "You have to specify what you want to drop.\n")
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

      session.User.CurrentWeight -= v.Weight
      Data.DB.Save(&session.User)

      Arg.WriteFull(session.Conn, "You've dropped " + v.Name + ".\n")

      return
    }
  }

  Arg.WriteFull(session.Conn, "You can't drop something that doesn't exist. Or can you?\n")
}

func NewDropCommand() CommandType {
  hc := NewCommand("drop", "'drop <name>' - Drop an item from your inventory.")
  hc.Handler = DropCommandHandler
  AllCommandsBig["drop"] =
  "Usage: 'drop <name>'\n" +
  "Drop an item from your inventory.\n"
  CommandsHelp[len(CommandsHelp)] = "drop"

  return hc
}
