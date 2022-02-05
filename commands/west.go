package Command

import (
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

func WestCommandHandler(cmd []string, session *Game.Session) {
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

  // Determine if room has NWestexit
  hasWestExit := strings.Contains(searchRoom.Exits, "West")

  // If so, move user and perform 'look' automatically
  if hasWestExit {
    // update position and save in database
    session.User.Y -= 1
    Data.DB.Save(&session.User)

    // output room 'look'
    session.Conn.Write([]byte(searchRoom.Name + "\n"))
    session.Conn.Write([]byte(searchRoom.Desc + "\n\n"))
    session.Conn.Write([]byte(Data.Cursor))

    return
  }

  // If not, show Data.InvalidMovement
  session.Conn.Write([]byte(Data.InvalidMovement))
  session.Conn.Write([]byte(Data.Cursor))
}

func NewWestCommand() CommandType {
  hc := NewCommand("West", "'West' - Move your character West.")
  hc.Handler = WestCommandHandler
  hc.Help =
  "West\n" +
  "Move your character West.\n"

  return hc
}
