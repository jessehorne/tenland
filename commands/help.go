package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
)

func HelpCommandHandler(cmd []string, session *Game.Session) {
  session.Conn.Write([]byte(Data.Help))
  session.Conn.Write([]byte(Data.Cursor))
}

func NewHelpCommand() CommandType {
  hc := NewCommand("help", "help")
  hc.Handler = HelpCommandHandler

  return hc
}
