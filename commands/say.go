package Command

import (
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/colors"
)

func SayCommandHandler(cmd []string, session *Game.Session) {
  // Make sure user is logged in
  if !session.Authed {
    session.Conn.Write([]byte("You have to be logged in to say things. Type 'help register' or 'help login'.\n"))
    session.Conn.Write([]byte(Data.Cursor))

    return
  }

  // Make sure len(cmd) > 1
  if len(cmd) < 2 {
    session.Conn.Write([]byte("What would you like to say? Try 'help say'.\n"))
    session.Conn.Write([]byte(Data.Cursor))

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
      sess.Conn.Write([]byte(packet))
      sess.Conn.Write([]byte(Data.Cursor))
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
