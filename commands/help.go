package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
)

func HelpCommandHandler(cmd []string, session *Game.Session) {
  // Check to see if user is asking for help for command
  if (len(cmd) > 1) {
    // Check to see if command exists
    if (AllCommands[cmd[1]] != "") {
      session.Conn.Write([]byte(AllCommandsBig[cmd[1]]))
    } else {
      session.Conn.Write([]byte("I'm sorry, that command doesn't exist. Try 'commands'."))
    }
  } else {
    session.Conn.Write([]byte(Data.Help))
  }

  session.Conn.Write([]byte("\n" + Data.Cursor))
}

func NewHelpCommand() CommandType {
  hc := NewCommand("help", "'help' - General helpful information.")
  hc.Handler = HelpCommandHandler
  AllCommandsBig["help"] =
  "Usage: 'help'\n" +
  "Provides a brief introduction to Tenland and how to play it.\n"

  return hc
}
