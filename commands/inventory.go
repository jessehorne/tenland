package Command

import (
  "strconv"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

func InventoryCommandHandler(cmd []string, session *Game.Session) {
  // Verify that user is logged in
  if !session.Authed {
    session.Conn.Write([]byte("You can't do this unless you're logged in.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  items := Model.ItemsGetByUserID(Data.DB, session.User.ID)

  message := "Inventory\n"

  if len(items) == 0 {
    message = message + "You are no items in your inventory.\n"
  } else {
    for i,v := range items {
      if message == "" {
        message = strconv.Itoa(i+1) + ". " + v.Name + "\n"
      } else {
        message = message + strconv.Itoa(i+1) + ". " + v.Name + "\n"
      }
    }
  }

  session.Conn.Write([]byte(message))

  session.Conn.Write([]byte("\n" + Data.Cursor))
}

func NewInventoryCommand() CommandType {
  hc := NewCommand("inventory", "'inventory' - Show your inventory.")
  hc.Handler = InventoryCommandHandler
  AllCommandsBig["inventory"] =
  "Usage: 'inventory'\n" +
  "Show your inventory.\n"
  CommandsHelp[len(CommandsHelp)] = "inventory"

  return hc
}
