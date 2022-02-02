package Command

import (
  "fmt"
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
)

func CoordsCommandHandler(cmd []string, session *Game.Session) {
  output := fmt.Sprintf("Coordinates (X, Y): (%d, %d)\n", session.X, session.Y)
  session.Conn.Write([]byte(output))
  session.Conn.Write([]byte(Data.Cursor))
}

func NewCoordsCommand() CommandType {
  c := NewCommand("coords", "'coords' - Shows current player coordinates.")
  c.Handler = CoordsCommandHandler
  c.Help =
  "coords\n" +
  "Shows your current coordinates.\n"

  return c
}
