package Arg

import (
  "net"
  "fmt"
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/commands"
  "github.com/jessehorne/tenland/game"
)

func Handle(n int, buf [512]byte, session *Game.Session) {
  // Split command
  cmd := string(buf[0:n-1])
  splitCmd := strings.Split(cmd, " ")

  if splitCmd[0] == "exit" {
    Write([]byte(Data.Goodbye), session.Conn)
    session.Conn.Close()
    fmt.Println("[USER DISCONNECTED]", session.IP)
  } else {
    f, found := Command.Run[splitCmd[0]]

    if found {
      f.(Command.CommandType).Handler.(func([]string, *Game.Session))(splitCmd, session)
    } else {
      WriteFull([]byte(Data.UnknownCommand), session.Conn)
    }
  }
}

func Write(buf []byte, conn net.Conn) {
  conn.Write(buf[0:])
  conn.Write([]byte("\n"))
}

func Cursor(conn net.Conn) {
  conn.Write([]byte(Data.Cursor))
}

func WriteFull(buf []byte, conn net.Conn) {
  Write(buf, conn)
  Cursor(conn)
}
