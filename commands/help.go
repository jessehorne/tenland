package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/arg"
)

func HelpCommandHandler(cmd []string, session *Game.Session) {
  // Check to see if user is asking for help for command
  if (len(cmd) > 1) {
    // Check to see if command exists
    if (AllCommands[cmd[1]] != "") {
      Arg.WriteFull(session.Conn, AllCommandsBig[cmd[1]] + "\n")
    } else {
      Arg.WriteFull(session.Conn, "I'm sorry, that command doesn't exist. Try 'commands'.\n")
    }
  } else {
    Arg.WriteFull(session.Conn, Data.Help + "\n")
  }
}

func NewHelpCommand() CommandType {
  hc := NewCommand("help", "'help' - General helpful information.")
  hc.Handler = HelpCommandHandler
  AllCommandsBig["help"] =
  "Usage: 'help'\n" +
  "Provides a brief introduction to Tenland and how to play it.\n"
  CommandsHelp[len(CommandsHelp)] = "help"

  return hc
}
