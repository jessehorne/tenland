package Command

import (
  "strconv"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/colors"
)

func WorthCommandHandler(cmd []string, session *Game.Session) {
  // Verify that user is logged in
  if !session.Authed {
    session.Conn.Write([]byte("You can't do this unless you're logged in.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  message := "" +
  Colors.Cyan("[Bank] ") + Colors.Yellow(strconv.Itoa(session.User.GoldBank)) + "\n" +
  Colors.Cyan("[Hand] ") + Colors.Yellow(strconv.Itoa(session.User.GoldHand)) + "\n"

  session.Conn.Write([]byte(message))
  session.Conn.Write([]byte(Data.Cursor))
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
