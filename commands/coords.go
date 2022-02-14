package Command

import (
  "fmt"
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
)

func CoordsCommandHandler(cmd []string, session *Game.Session) {
  output := fmt.Sprintf("Coordinates (X, Y): (%d, %d)\n", session.User.X, session.User.Y)
  session.Conn.Write([]byte(output))
  session.Conn.Write([]byte(Data.Cursor))
}

func NewCoordsCommand() CommandType {
  c := NewCommand("coords", "'coords' - Shows current player coordinates.")
  c.Handler = CoordsCommandHandler
  AllCommandsBig["coords"] =
  "Usage: 'coords'\n" +
  "Shows your current coordinates.\n"
  CommandsHelp[len(CommandsHelp)] = "coords"

  return c
}
