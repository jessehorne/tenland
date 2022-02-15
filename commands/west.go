package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/arg"
)

func WestCommandHandler(cmd []string, session *Game.Session) {
  // Get users current position
  x := session.User.X - 1
  y := session.User.Y

  // Get room from DB at current coords
  searchRoom := Model.Room{}
  result := Data.DB.Where(Model.Room{X: x, Y: y}).First(&searchRoom)

  if result.RowsAffected == 0 {
    // This runs if there is no room where the user is going
    if !session.User.IsBuilder {
      // Can't move because the user isn't even in a room
      Arg.WriteFull(session.Conn, "You can't enter a room from the abyss.\n")

      return
    } else {
      session.User.X = x
      session.User.Y = y

      Data.DB.Save(&session.User)

      Arg.WriteFull(session.Conn, "You went west into the abyss.\n")
    }
  } else {
    // If there is a room, then move there
    session.User.X = x
    session.User.Y = y

    Data.DB.Save(&session.User)

    Arg.Write(session.Conn, searchRoom.Title + "\n")
    Arg.WriteFull(session.Conn, searchRoom.Desc + "\n\n")
  }

}

func NewWestCommand() CommandType {
  hc := NewCommand("west", "'west' - Move your character west.")
  hc.Handler = WestCommandHandler
  AllCommandsBig["west"] =
  "Usage: 'west'\n" +
  "Move your character west.\n"

  CommandsHelp[len(CommandsHelp)] = "west"

  return hc
}
