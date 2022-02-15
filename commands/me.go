package Command

import (
  "fmt"

  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/arg"
)

func MeCommandHandler(cmd []string, session *Game.Session) {
  // Make sure user is logged in
  if !session.Authed {
    Arg.WriteFull(session.Conn, "You have to be logged in to say things. Type 'help register' or 'help login'.\n")

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

  Arg.WriteFull(session.Conn, fmt.Sprintf(card,
    session.User.Username,
    session.User.GoldBank,
    session.User.GoldHand,
    session.User.ST,
    session.User.GetMaxCarryWeight(),
    session.User.CurrentWeight))
}

func NewMeCommand() CommandType {
  hc := NewCommand("me", "'me' - Detailed information on your character.")
  hc.Handler = MeCommandHandler
  AllCommandsBig["me"] =
  "Usage: 'me'\n" +
  "Detailed information on your character.\n"

  CommandsHelp[len(CommandsHelp)] = "me"

  return hc
}
