package Command

import (
  "strings"

  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/colors"
  "github.com/jessehorne/tenland/arg"
)

func SayCommandHandler(cmd []string, session *Game.Session) {
  // Make sure user is logged in
  if !session.Authed {
    Arg.WriteFull(session.Conn, "You have to be logged in to say things. Type 'help register' or 'help login'.\n")

    return
  }

  // Make sure len(cmd) > 1
  if len(cmd) < 2 {
    Arg.WriteFull(session.Conn, "What would you like to say? Try 'help say'.\n")

    return
  }

  // Get users current position
  x := session.User.X
  y := session.User.Y

  // Get message from cmd
  message := strings.Join(cmd[1:], " ")

  // Build packet
  packet := "\n" + Colors.Green("[Local]") + Colors.Blue(" (player) ") + Colors.Cyan(session.User.Username + ": " + message + "\n")

  // Send message to all users in this area
  for _,sess := range Game.Sessions {
    if sess.User.X == x && sess.User.Y == y {
      Arg.WriteFull(sess.Conn, packet)
    }
  }

}

func NewSayCommand() CommandType {
  hc := NewCommand("say", "'say <message>' - Say something outloud that can be heard in the room you're currently in.")
  hc.Handler = SayCommandHandler
  AllCommandsBig["say"] =
  "Usage: 'say <message>'\n" +
  "Say something outloud that can be heard in the room you're currently in.\n"

  CommandsHelp[len(CommandsHelp)] = "say"

  return hc
}
