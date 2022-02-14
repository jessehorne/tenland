package Command

import (
  "strconv"
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
)

func CommandsCommandHandler(cmd []string, session *Game.Session) {
  session.Conn.Write([]byte("[Commands]\n"))

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

    session.Conn.Write([]byte(msg))
  }

  // Send cursor
  session.Conn.Write([]byte(Data.Cursor))
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
