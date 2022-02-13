package Command

import (
  "fmt"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
)

func MeCommandHandler(cmd []string, session *Game.Session) {
  // Make sure user is logged in
  if !session.Authed {
    session.Conn.Write([]byte("You have to be logged in to say things. Type 'help register' or 'help login'.\n"))
    session.Conn.Write([]byte(Data.Cursor))

    return
  }

  card := `
Name: %s

Wealth
======
Gold in Bank: %d
Gold in Hand: %d

Stats
=====
Strength: %d
Max Carry Weight: %dkg              Carrying: %.2fkg
`

  session.Conn.Write([]byte(fmt.Sprintf(card,
    session.User.Username,
    session.User.GoldBank,
    session.User.GoldHand,
    session.User.ST,
    session.User.GetMaxCarryWeight(),
    session.User.CurrentWeight)))
  session.Conn.Write([]byte("\n" + Data.Cursor))
}

func NewMeCommand() CommandType {
  hc := NewCommand("me", "'me' - Detailed information on your character.")
  hc.Handler = MeCommandHandler
  AllCommandsBig["me"] =
  "Usage: 'me'\n" +
  "Detailed information on your character.\n"

  return hc
}
