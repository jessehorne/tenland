package Command

import (
  "net"

  "github.com/jessehorne/tenland/data"
)

func HelpCommandHandler(cmd []string, conn net.Conn) {
  conn.Write([]byte(Data.Help))
  conn.Write([]byte(Data.Cursor))
}

func NewHelpCommand() CommandType {
  hc := NewCommand("help", "help")
  hc.Handler = HelpCommandHandler

  return hc
}
