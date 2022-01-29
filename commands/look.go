package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

func LookCommandHandler(cmd []string, session *Game.Session) {
  // Get users current position
  x := session.X
  y := session.Y

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

  // The room was found, print title, desc + exits
  session.Conn.Write([]byte(searchRoom.Name + "\n"))
  session.Conn.Write([]byte(searchRoom.Desc + "\n\n"))
  session.Conn.Write([]byte(Data.Cursor))
}

func NewLookCommand() CommandType {
  c := NewCommand("look", "'look' - Get visual and auditory information on current room.")
  c.Handler = LookCommandHandler

  return c
}
