package Command

import (
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

func NorthCommandHandler(cmd []string, session *Game.Session) {
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

  // Determine if room has North exit
  hasNorthExit := strings.Contains(searchRoom.Exits, "north")

  // If so, move user north and perform 'look' automatically
  if hasNorthExit {
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

func NewNorthCommand() CommandType {
  hc := NewCommand("north", "'north' - Move your character north.")
  hc.Handler = NorthCommandHandler
  hc.Help =
  "north\n" +
  "Move your character north.\n"

  return hc
}
