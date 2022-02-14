package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
)

func BeginCommandHandler(cmd []string, session *Game.Session) {
  session.Conn.Write([]byte(Data.Begin))
  session.Conn.Write([]byte(Data.Cursor))
}

func NewBeginCommand() CommandType {
  bc := NewCommand("begin", "'begin' - Tells you how to start playing the game.")
  bc.Handler = BeginCommandHandler
  AllCommandsBig["begin"] =
  "Usage: 'begin'\n" +
  "Gives you a brief explanation on how to begin if you're new to Tenland.\n"
  CommandsHelp[len(CommandsHelp)] = "begin"

  return bc
}
