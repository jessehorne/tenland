package Command

import (
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

func EastCommandHandler(cmd []string, session *Game.Session) {
  // Get users current position
  x := session.User.X
  y := session.User.Y

  // Get room from DB at current coords
  searchRoom := Model.Room{}
  result := Data.DB.Where(Model.Room{X: x, Y: y}).First(&searchRoom)

  if result.RowsAffected == 0 {
    // Can't move because the user isn't even in a room
    session.Conn.Write([]byte("You can't enter a room from the abyss.\n"))
    session.Conn.Write([]byte(Data.Cursor))

    return
  }

  // Determine if room has NEastexit
  hasEastExit := strings.Contains(searchRoom.Exits, "east")

  // If so, move user and perform 'look' automatically
  if hasEastExit {
    // update position and save in database
    session.User.Y -= 1
    Data.DB.Save(&session.User)

    // output room 'look'
    session.Conn.Write([]byte(searchRoom.Title + "\n"))
    session.Conn.Write([]byte(searchRoom.Desc + "\n\n"))
    session.Conn.Write([]byte(Data.Cursor))

    return
  }

  // If not, show Data.InvalidMovement
  session.Conn.Write([]byte(Data.InvalidMovement))
  session.Conn.Write([]byte(Data.Cursor))
}

func NewEastCommand() CommandType {
  hc := NewCommand("east", "'east' - Move your character east.")
  hc.Handler = EastCommandHandler
  hc.Help =
  "east\n" +
  "Move your character east.\n"

  return hc
}
