package Game

import (
  "net"

  "github.com/jessehorne/tenland/models"
)

type Session struct {
  User Model.User
  Authed bool
  IP string
  Conn net.Conn
}

func NewSession(conn net.Conn) Session {
  newSession := Session{
    Authed: false,
    IP: conn.LocalAddr().String(),
    Conn: conn,
  }

  return newSession
}
