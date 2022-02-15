package Command

import (
  "strconv"

  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/colors"
  "github.com/jessehorne/tenland/arg"
)

func WorthCommandHandler(cmd []string, session *Game.Session) {
  // Verify that user is logged in
  if !session.Authed {
    Arg.WriteFull(session.Conn, "You can't do this unless you're logged in.\n")
    return
  }

  message := "" +
  Colors.Cyan("[Bank] ") + Colors.Yellow(strconv.Itoa(session.User.GoldBank)) + "\n" +
  Colors.Cyan("[Hand] ") + Colors.Yellow(strconv.Itoa(session.User.GoldHand)) + "\n"

  Arg.WriteFull(session.Conn, message)
}

func NewWorthCommand() CommandType {
  hc := NewCommand("worth", "'worth' - Display information about your worth.")
  hc.Handler = WorthCommandHandler
  AllCommandsBig["worth"] =
  "Usage: 'worth'\n" +
  "Display information how how much gold you're carrying and how much is in your bank.\n"

  CommandsHelp[len(CommandsHelp)] = "worth"

  return hc
}
