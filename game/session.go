package Game

import (
  "net"

  "github.com/thanhpk/randstr"

  "github.com/jessehorne/tenland/models"
)

type Session struct {
  User Model.User
  Authed bool
  IP string
  Conn net.Conn
  ID string
}

var Sessions = map[string]Session{}

func NewSession(conn net.Conn) Session {
  unique := randstr.Hex(16)

  newSession := Session{
    Authed: false,
    IP: conn.LocalAddr().String(),
    Conn: conn,
    ID: unique,
  }

  Sessions[unique] = newSession

  return newSession
}
