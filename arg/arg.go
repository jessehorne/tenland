package Arg

import (
  "net"
  "fmt"

  "github.com/jessehorne/tenland/data"
)

func Write(conn net.Conn, buf string) {
  i, err := conn.Write([]byte(buf))

  if err != nil {
    fmt.Println("[ERROR] arg.go func Write(...)", i)
  }
}

func Cursor(conn net.Conn) {
  i, err := conn.Write([]byte(Data.Cursor))

  if err != nil {
    fmt.Println("[ERROR] arg.go func Cursor(...)", i)
  }
}

func WriteFull(conn net.Conn, buf string) {
  Write(conn, buf)
  Cursor(conn)
}
