package Command

import (
  "strings"

  // "github.com/go-playground/validator/v10"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/arg"
)


// The main handler for the 'build' command.
// Examples of use...
// Create a room: 'build create'
// Change the title of the current room: 'build set title An Example Room Name'
// Change the description of the current room: 'build set desc An example room description!'
// Add an exit: 'build exit add north'
// Remove an exit: 'build exit remove north'
func BuildCommandHandler(cmd []string, session *Game.Session) {
  // Auth
    // Verify that user is logged in
    if !session.Authed {
      Arg.WriteFull(session.Conn, "You can't do this unless you're logged in.\n")
      return
    }

    // Verify that user is a builder
    if !session.User.IsBuilder {
      Arg.WriteFull(session.Conn, "You're not a builder!\n")
      return
    }

  // Validate that cmd is at least 2 in length
  if len(cmd) < 2 {
    Arg.WriteFull(session.Conn, "Invalid. You must specify a command. Please type 'help build'.\n")
    return
  }

  first := cmd[1]

  if first == "set" {
    SetRoom(cmd[1:], session)
    return
  } else if first == "create" {
    CreateRoom(session)
    return
  } else if first == "exit" {
    ExitRoom(cmd[1:], session)
    return
  }

  Arg.WriteFull(session.Conn, "Invalid. Invalid command. Please type 'help build'.\n")
}

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
    Arg.WriteFull(session.Conn, "Invalid. A room already exists here. Type 'look'.\n")

    return
  } else {
    // If it doesn't exist, create it
    _, err := Model.NewRoom(Data.DB, x, y, "Construction Zone", "This room is new and unfinished.")

    if err != nil {
      Arg.WriteFull(session.Conn, "Sorry! Something went wrong creating the room.\n")
    } else {
      Arg.WriteFull(session.Conn, "You've created a room!\n")
    }
  }
}

// Manage room exits
// Examples...
// build exit add north
// build exit remove north
func ExitRoom(cmd []string, session *Game.Session) {
  // len(cmd) should be at least 3
  if len(cmd) < 3 {
    Arg.WriteFull(session.Conn, "Invalid use of 'exit'. Missing arguments. Please type 'help build'.\n")
    return
  }

  which := cmd[1] // add or remove
  direction := cmd[2] // north, south, east, west, etc

  // Get users current position
  x := session.User.X
  y := session.User.Y

  // Get room at current user position
  searchRoom := Model.Room{}
  result := Data.DB.Where(Model.Room{
    X: x,
    Y: y,
  }).First(&searchRoom)

  if result.RowsAffected == 0 {
    Arg.WriteFull(session.Conn, "Invalid use of 'exit'. You can't set exits in the abyss. Please type 'help build'.\n")
    return
  }

  exits := strings.Split(searchRoom.Exits, ",")

  if which == "add" {
    // Loop over exits, if direction not found, then you can add it
    found := false
    for _, value := range exits {
      if value == direction {
        found = true
        break
      }
    }

    if found {
      Arg.WriteFull(session.Conn, "Invalid use of 'exit'. That exit already exists. Please type 'help build'.\n")
      return
    } else {
      // add exit and save
      if searchRoom.Exits == "" {
        searchRoom.Exits = direction
      } else {
        searchRoom.Exits = searchRoom.Exits + "," + direction
      }

      Data.DB.Save(&searchRoom)

      Arg.WriteFull(session.Conn, "You've added an exit to the current room!\n")

      return
    }
  } else if which == "remove" {
    // Iterate over current room exits. If found, you can delete it
    found := false
    index := -1
    for i, value := range exits {
      if value == direction {
        found = true
        index = i
      }
    }

    if found {
      // delete exit
      newExits := []string{}

      newExits = append(newExits, exits[:index]...)
      newExits = append(newExits, exits[index+1:]...)

      searchRoom.Exits = strings.Join(newExits, ",")

      Data.DB.Save(&searchRoom)

      Arg.WriteFull(session.Conn, "You've removed an exit from the current room!\n")

      return
    } else {
      Arg.WriteFull(session.Conn, "Invalid use of 'exit'. That exit doesn't exist in this room. Please type 'help build'.\n")
      return
    }

  }

  // If it gets here, means the user typed wrong command (not 'add' nor 'remove')
  Arg.WriteFull(session.Conn, "Invalid use of 'exit'. Invalid option to 'exit'. Please type 'help build'.\n")
}

// Used to set room variables such as 'title' and 'desc'
func SetRoom(cmd []string, session *Game.Session) {
  // len(cmd) should be at least 3 (set, title OR desc, data) example: 'set title A Small Room'
  if len(cmd) < 3 {
    Arg.WriteFull(session.Conn, "Invalid use of 'set'. Missing arguments. Please type 'help build'.\n")
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
    Arg.WriteFull(session.Conn, "Invalid use of 'set'. You can't build in the abyss. Please type 'help build'.\n")
    return
  }

  data := strings.Join(cmd[2:], " ")
  which := cmd[1]

  if which == "title" {
    searchRoom.Title = data
    Data.DB.Save(&searchRoom)

    Arg.WriteFull(session.Conn, "You've updated the room <title>.\n")
    return
  } else if which == "desc" {
    searchRoom.Desc = data
    Data.DB.Save(&searchRoom)

    Arg.WriteFull(session.Conn, "You've updated the room <description>.\n")
    return
  }

  Arg.WriteFull(session.Conn, "Invalid use of 'set'. Invalid arguments. Please type 'help build'.\n")
}


func NewBuildCommand() CommandType {
  hc := NewCommand("build", "'build <command> <command_arg> <data>' - Tool for builders.")
  hc.Handler = BuildCommandHandler
  AllCommandsBig["build"] =
  "Usage: 'build <command> <args...>\n" +
  "Lets builders edit the area description.\n" +
  "Examples:\n" +
  "build create\n" +
  "build set title A Small Room\n" +
  "build set desc A description of a small room.\n" +
  "build exit add north\n" +
  "build exit remove north\n"
  CommandsHelp[len(CommandsHelp)] = "build"

  return hc
}
