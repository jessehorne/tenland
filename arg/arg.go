package Arg

import (
  "net"
  "fmt"
  "github.com/jessehorne/tenland/data"
)

func Handle(n int, buf [512]byte, conn net.Conn) {
  cmd := string(buf[0:n-1])
  if cmd == "exit" {
    Write([]byte(Data.Goodbye), conn)
    conn.Close()
    fmt.Println("[USER DISCONNECTED]", conn.LocalAddr().String())
  } else if cmd == "help" {
    WriteFull([]byte(Data.Help), conn)
  } else {
    WriteFull([]byte(Data.UnknownCommand), conn)
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
