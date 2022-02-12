package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/colors"
)

func PlayersCommandHandler(cmd []string, session *Game.Session) {
  // loop through all sessions and build a string of connected users
  users := Colors.Cyan("Online Users\n\n")
  userCount := 0

  for _,v := range Game.Sessions {
    users += v.User.Username + "\n"
    userCount += 1
  }

  if userCount == 0 {
    users += "No users online...Weird, right?"
  }

  session.Conn.Write([]byte(users))

  session.Conn.Write([]byte("\n" + Data.Cursor))
}

func NewPlayersCommand() CommandType {
  hc := NewCommand("players", "'players' - Get a list of all online players.")
  hc.Handler = PlayersCommandHandler
  AllCommandsBig["players"] =
  "Usage: 'players'\n" +
  "Get a list of all online players.\n"

  return hc
}
