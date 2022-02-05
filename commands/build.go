package Command

import (
  "strings"

  // "github.com/go-playground/validator/v10"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)


// Creates blank room at users current position
func CreateRoom(session *Game.Session) {
  // Get user coordinates
  x := session.User.X
  y := session.User.Y

  // Check if room exists
  searchRoom := Model.Room{}
  result := Data.DB.Where(Model.Room{X: x, Y: y}).First(&searchRoom)

  // If it exists, update it
  if result.RowsAffected > 0 {
    session.Conn.Write([]byte("Invalid. A room already exists here. Type 'look'.\n"))
    session.Conn.Write([]byte(Data.Cursor))

    return
  } else {
    // If it doesn't exist, create it
    _, err := Model.NewRoom(Data.DB, x, y, "Construction Zone", "This room was just created and has no exits.")

    if err != nil {
      session.Conn.Write([]byte("Sorry! Something went wrong creating the room.\n"))
      session.Conn.Write([]byte(Data.Cursor))
    } else {
      session.Conn.Write([]byte("You've created a room!\n"))
      session.Conn.Write([]byte(Data.Cursor))
    }
  }
}

// Used to set room variables such as 'title' and 'desc'
func SetRoom(cmd []string, session *Game.Session) {
  // len(cmd) should be at least 3 (set, title OR desc, data) example: 'set title A Small Room'
  if len(cmd) < 3 {
    session.Conn.Write([]byte("Invalid use of 'set'. Missing arguments. Please type 'help build'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  // Validate that user is in room

  // Get user coordinates
  x := session.User.X
  y := session.User.Y

  // Check if room exists
  searchRoom := Model.Room{}
  result := Data.DB.Where(Model.Room{X: x, Y: y}).First(&searchRoom)

  if result.RowsAffected == 0 {
    session.Conn.Write([]byte("Invalid use of 'set'. You can't build in the abyss. Please type 'help build'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  data := strings.Join(cmd[2:], " ")
  which := cmd[1]

  if which == "title" {
    searchRoom.Title = data
    Data.DB.Save(&searchRoom)

    session.Conn.Write([]byte("You've updated the room <title>.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  } else if which == "desc" {
    searchRoom.Desc = data
    Data.DB.Save(&searchRoom)

    session.Conn.Write([]byte("You've updated the room <description>.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  session.Conn.Write([]byte("Invalid use of 'set'. Invalid arguments. Please type 'help build'.\n"))
  session.Conn.Write([]byte(Data.Cursor))
}


func BuildCommandHandler(cmd []string, session *Game.Session) {
  // Auth
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

  // Validate that cmd is at least 2 in length
  if len(cmd) < 2 {
    session.Conn.Write([]byte("Invalid. You must specify a command. Please type 'help build'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  first := cmd[1]

  if first == "set" {
    SetRoom(cmd[1:], session)

    return
  } else if first == "create" {
    CreateRoom(session)

    return
  }

  session.Conn.Write([]byte("Invalid. Invalid command. Please type 'help build'.\n"))
  session.Conn.Write([]byte(Data.Cursor))
}


func NewBuildCommand() CommandType {
  hc := NewCommand("build", "'build <command> <command_arg> <data>' - Tool for builders.")
  hc.Handler = BuildCommandHandler
  hc.Help =
  "build\n" +
  "Lets builders edit the area description.\n" +
  "Usage:\n" +
  "build set title A Small Room\n" +
  "build set desc A description of a small room.\n"

  return hc
}
