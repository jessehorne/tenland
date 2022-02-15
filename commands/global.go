package Command

import (
  "strings"

  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/colors"
  "github.com/jessehorne/tenland/arg"
)

func GlobalCommandHandler(cmd []string, session *Game.Session) {
  // Make sure user is logged in
  if !session.Authed {
    Arg.WriteFull(session.Conn, "You have to be logged in to send messages to Global.\n")

    return
  }

  // Make sure len(cmd) > 1
  if len(cmd) < 2 {
    Arg.WriteFull(session.Conn, "What would you like to send to Global? Try 'help global'.\n")

    return
  }

  // Get message from cmd
  message := strings.Join(cmd[1:], " ")

  // Build packet
  packet := "\n" + Colors.Yellow("[Global]") + Colors.Blue(" (player) ") + Colors.Cyan(session.User.Username + ": " + message + "\n")

  // Loop through all sessions and send message
  for _,val := range Game.Sessions {
    Arg.WriteFull(val.Conn, packet)
  }
}

func NewGlobalCommand() CommandType {
  gc := NewCommand("global", "'global <message>' - Send message to Global channel, which everyone can see.")
  gc.Handler = GlobalCommandHandler
  AllCommandsBig["global"] =
  "Usage: 'global <message>'\n" +
  "Send message to Global channel, which everyone can see.\n" +
  "Example...\n" +
  "'global Hello everyone!'"
  CommandsHelp[len(CommandsHelp)] = "global"

  return gc
}
