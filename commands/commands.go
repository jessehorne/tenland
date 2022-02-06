package Command

import (
  "strconv"
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
)

func CommandsCommandHandler(cmd []string, session *Game.Session) {
  session.Conn.Write([]byte("[Commands]\n"))

  // Loop through commands and send each line
  i := 1
  for key, element := range AllCommands {
    session.Conn.Write([]byte(strconv.Itoa(i) + ". " + key + "\t\t\t" + element + "\n"))
    i++
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

  return c
}
