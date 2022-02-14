package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

func NorthCommandHandler(cmd []string, session *Game.Session) {
  // Get users current position
  x := session.User.X
  y := session.User.Y - 1

  // Get room from DB at current coords
  searchRoom := Model.Room{}
  result := Data.DB.Where(Model.Room{X: x, Y: y}).First(&searchRoom)

  if result.RowsAffected == 0 {
    // This runs if there is no room where the user is going
    if !session.User.IsBuilder {
      // Can't move because the user isn't even in a room
      session.Conn.Write([]byte("You can't enter a room from the abyss.\n"))
      session.Conn.Write([]byte(Data.Cursor))

      return
    } else {
      session.User.X = x
      session.User.Y = y

      Data.DB.Save(&session.User)

      session.Conn.Write([]byte("You went north into the abyss.\n"))
      session.Conn.Write([]byte(Data.Cursor))
    }
  } else {
    // If there is a room, then move there
    session.User.X = x
    session.User.Y = y

    Data.DB.Save(&session.User)

    session.Conn.Write([]byte(searchRoom.Title + "\n"))
    session.Conn.Write([]byte(searchRoom.Desc + "\n\n"))
    session.Conn.Write([]byte(Data.Cursor))
  }

}

func NewNorthCommand() CommandType {
  hc := NewCommand("north", "'north' - Move your character north.")
  hc.Handler = NorthCommandHandler
  AllCommandsBig["north"] =
  "Usage: 'north'\n" +
  "Move your character north.\n"

  CommandsHelp[len(CommandsHelp)] = "north"

  return hc
}
