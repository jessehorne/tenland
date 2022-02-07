package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/colors"
)

func LookCommandHandler(cmd []string, session *Game.Session) {
  // Make sure user is logged in
  if !session.Authed {
    session.Conn.Write([]byte("You see nothing. You have to be logged in to look around. Type 'help login' or 'help register'.\n"))
    session.Conn.Write([]byte(Data.Cursor))

    return
  }

  // Get users current position
  x := session.User.X
  y := session.User.Y

  // Get room at current user position
  searchRoom := Model.Room{}
  result := Data.DB.Where(Model.Room{
    X: x,
    Y: y,
  }).First(&searchRoom)

  // Handle if room doesn't exist in database
  if result.RowsAffected == 0 {
    session.Conn.Write([]byte(Data.Abyss))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  // Get list of users in room
  users := ""

  for _, u := range Game.Sessions {
    if u.User.X == x && u.User.Y == y && u.ID != session.ID {
      if users == "" {
        users = u.User.Username
      } else {
        users = users + ", " + u.User.Username
      }
    }
  }

  if users == "" {
    users = "There doesn't appear to be anyone here."
  }

  // The room was found, print title, desc + exits
  session.Conn.Write([]byte(searchRoom.Title + "\n"))
  session.Conn.Write([]byte(searchRoom.Desc + "\n\n"))
  session.Conn.Write([]byte(Colors.Yellow("Users: " + users + "\n")))
  session.Conn.Write([]byte("Exits: [" + searchRoom.Exits + "]\n\n"))

  session.Conn.Write([]byte(Data.Cursor))
}

func NewLookCommand() CommandType {
  c := NewCommand("look", "'look' - Get visual and auditory information on current room.")
  c.Handler = LookCommandHandler
  AllCommandsBig["look"] =
  "Usage: 'look'\n" +
  "Gives you a detailed description of the area you're current standing in.\n"

  return c
}
