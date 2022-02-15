package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/arg"
)

func TakeCommandHandler(cmd []string, session *Game.Session) {
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
    Arg.WriteFull(session.Conn, "There is no such item in the abyss.\n")
    return
  }

  // Make sure argument is provided
  if len(cmd) < 2 {
    Arg.WriteFull(session.Conn, "You have to specify what you want to take.\n")
    return
  }

  // Get list of items in room
  allItems := Model.ItemsGetByRoom(Data.DB, x, y)

  for _,v := range allItems {
    if v.Name == cmd[1] && !v.Held {
      if float32(session.User.GetMaxCarryWeight()) >= session.User.CurrentWeight + v.Weight {
        // Item exists, now let user pick it up
        v.Held = true
        v.UserID = session.User.ID
        Data.DB.Save(&v)

        session.User.CurrentWeight += v.Weight
        Data.DB.Save(&session.User)

        Arg.WriteFull(session.Conn, "You've picked up " + v.Name + " and put it in your inventory.\n")

        return
      } else {
        Arg.WriteFull(session.Conn, "You can't pick that up. It's too heavy!\n")

        return
      }
    }
  }

  Arg.WriteFull(session.Conn, "There doesn't appear to be anything like that here to pick up.\n")
}

func NewTakeCommand() CommandType {
  hc := NewCommand("take", "'take <name>' - Pick up an item and put it in your inventory.")
  hc.Handler = TakeCommandHandler
  AllCommandsBig["take"] =
  "Usage: 'take <name>'\n" +
  "Pick up an item and put it in your inventory.\n"

  CommandsHelp[len(CommandsHelp)] = "take"

  return hc
}
