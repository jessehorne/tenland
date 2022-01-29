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
  c := NewCommand("coords", "coords")
  c.Handler = CoordsCommandHandler

  return c
}
