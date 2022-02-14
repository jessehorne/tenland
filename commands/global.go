package Command

import (
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/colors"
)

func GlobalCommandHandler(cmd []string, session *Game.Session) {
  // Make sure user is logged in
  if !session.Authed {
    session.Conn.Write([]byte("You have to be logged in to send messages to Global.\n"))
    session.Conn.Write([]byte(Data.Cursor))

    return
  }

  // Make sure len(cmd) > 1
  if len(cmd) < 2 {
    session.Conn.Write([]byte("What would you like to send to Global? Try 'help global'.\n"))
    session.Conn.Write([]byte(Data.Cursor))

    return
  }

  // Get message from cmd
  message := strings.Join(cmd[1:], " ")

  // Build packet
  packet := "\n" + Colors.Yellow("[Global]") + Colors.Blue(" (player) ") + Colors.Cyan(session.User.Username + ": " + message + "\n")

  // Loop through all sessions and send message
  for _,val := range Game.Sessions {
    val.Conn.Write([]byte(packet))
    val.Conn.Write([]byte(Data.Cursor))
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
