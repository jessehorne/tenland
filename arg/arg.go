package Arg

import (
  "net"
  "fmt"
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/commands"
)

func Handle(n int, buf [512]byte, conn net.Conn) {
  // Split command
  cmd := string(buf[0:n-1])
  splitCmd := strings.Split(cmd, " ")

  if splitCmd[0] == "exit" {
    Write([]byte(Data.Goodbye), conn)
    conn.Close()
    fmt.Println("[USER DISCONNECTED]", conn.LocalAddr().String())
  } else {
    f, found := Command.Run[splitCmd[0]]

    if found {
      f.(Command.CommandType).Handler.(func([]string, net.Conn))(splitCmd, conn)
    } else {
      WriteFull([]byte(Data.UnknownCommand), conn)
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
