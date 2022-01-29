package Game

import (
  "net"

  "github.com/jessehorne/tenland/models"
)

type Session struct {
  User Model.User
  Authed bool
  Admin bool
  IP string
  Conn net.Conn
  X int
  Y int
}

func NewSession(conn net.Conn) Session {
  newSession := Session{
    Authed: false,
    Admin: false,
    IP: conn.LocalAddr().String(),
    Conn: conn,
    X: 0,
    Y: 0,
  }

  return newSession
}
