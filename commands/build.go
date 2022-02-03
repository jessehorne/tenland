package Command

import (
  "strings"

  "github.com/go-playground/validator/v10"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

type BuildInput struct {
  Title string `validate:"required,lte=255"`
  Desc string `validate:"required"`
}

func BuildCommandHandler(cmd []string, session *Game.Session) {
  // Verify that the command has <title> and <desc>
  if len(cmd) < 3 {
    session.Conn.Write([]byte("You have to supply a <title> AND <description>.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  // Verify that user is logged in
  if !session.Authed {
    session.Conn.Write([]byte("You can't do this unless you're logged in.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  // Verify that user is a builder
  if !session.User.IsBuilder {
    session.Conn.Write([]byte("You're not a builder!\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  // Verify input
  title := cmd[1]
  desc := strings.Join(cmd[2:], " ")

  validate := validator.New()

  input := &BuildInput{
    Title: title,
    Desc: desc,
  }

  err := validate.Struct(input)

  if err != nil {
    session.Conn.Write([]byte("Error! Be sure to use the following syntax...\n"))
    session.Conn.Write([]byte("build <title> <description>\n"))
    session.Conn.Write([]byte(Data.Cursor))
  }

  // Input is valid

  // Get user coordinates
  x := session.User.X
  y := session.User.Y

  // Check if room exists
  searchRoom := Model.Room{}
  result := Data.DB.Where(Model.Room{X: x, Y: y}).First(&searchRoom)

  // If it exists, update it
  if result.RowsAffected > 0 {
    searchRoom.Name = title
    searchRoom.Desc = desc

    Data.DB.Save(&searchRoom)

    session.Conn.Write([]byte("You've edited the current room!\n"))
    session.Conn.Write([]byte(Data.Cursor))
  } else {
    // If it doesn't exist, create it
    _, err := Model.NewRoom(Data.DB, x, y, title, desc)

    if err != nil {
      session.Conn.Write([]byte("Sorry! Something went wrong creating the room.\n"))
      session.Conn.Write([]byte(Data.Cursor))
    } else {
      session.Conn.Write([]byte("You've created a room!\n"))
      session.Conn.Write([]byte(Data.Cursor))
    }
  }

}

func NewBuildCommand() CommandType {
  hc := NewCommand("build", "'build <title> <description>' - Lets builders edit the area description.")
  hc.Handler = BuildCommandHandler
  hc.Help =
  "build\n" +
  "Lets builders edit the area description.\n"

  return hc
}
