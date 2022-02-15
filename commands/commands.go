package Command

import (
  "strconv"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/arg"
)

func CommandsCommandHandler(cmd []string, session *Game.Session) {
  Arg.Write(session.Conn, "[Commands]\n")

  for i := 0; i < len(CommandsHelp); i++ {
    msg := strconv.Itoa(i+1) + ". " + CommandsHelp[i]

    // Handles spaces (pls find better way)
    spaceCount := 20 - len(CommandsHelp[i])

    if i > 9 {
      spaceCount--
    }

    for x := 0; x < spaceCount; x++ {
      msg = msg + " "
    }

    msg = msg + AllCommands[CommandsHelp[i]] + "\n"

    Arg.Write(session.Conn, msg)
  }

  // Send cursor
  Arg.Cursor(session.Conn)
}

func NewCommandsCommand() CommandType {
  c := NewCommand("commands", "'commands' - Shows a list of all available commands.")
  c.Handler = CommandsCommandHandler
  AllCommandsBig["commands"] =
  "Usage: 'commands'\n" +
  "Displays a list of all possible commands.\n"
  CommandsHelp[len(CommandsHelp)] = "commands"

  return c
}
