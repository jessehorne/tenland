package Command

import (
  "fmt"
  "strconv"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/colors"
  "github.com/jessehorne/tenland/arg"
)

func LookAtItem(cmd []string, session *Game.Session) {
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
    Arg.WriteFull(session.Conn, "There is no such item in the abyss.\n")
    return
  }


  // Get list of items in room
  allItems := Model.ItemsGetByRoom(Data.DB, x, y)

  for _,v := range allItems {
    if v.Name == cmd[1] {
      Arg.Write(session.Conn, fmt.Sprintf("You look closely at %s (%.2fkg)...\n", v.Name, v.Weight))
      Arg.WriteFull(session.Conn, v.Description + "\n")

      return
    }
  }

  Arg.WriteFull(session.Conn, "There doesn't appear to be anything like that here.\n")

}

func LookInRoom(cmd []string, session *Game.Session) {
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
    Arg.WriteFull(session.Conn, Data.Abyss)
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

  // Get list of items in room
  items := ""

  allItems := Model.ItemsGetByRoom(Data.DB, x, y)

  for _,v := range allItems {
    if v.Held {
      continue
    }

    if items == "" {
      items = v.Name

      // if user is builder, add ID to item
      id := strconv.Itoa(int(v.ID))
      items = items + "(" + id + ")"
    } else {
      items = items + ", " + v.Name

      // if user is builder, add ID to item
      id := strconv.Itoa(int(v.ID))
      items = items + "(" + id + ")"
    }
  }

  if items == "" {
    items = "There doesn't appear to be any items here."
  }

  // The room was found, print title, desc + exits
  Arg.Write(session.Conn, searchRoom.Title + "\n")
  Arg.Write(session.Conn, searchRoom.Desc + "\n\n")
  Arg.Write(session.Conn, Colors.Yellow("Users: " + users + "\n"))
  Arg.WriteFull(session.Conn, "Items: " + items + "\n")
}

func LookCommandHandler(cmd []string, session *Game.Session) {
  // Make sure user is logged in
  if !session.Authed {
    Arg.WriteFull(session.Conn, "You see nothing. You have to be logged in to look around. Type 'help login' or 'help register'.\n")

    return
  }

  // Determine if looking generally or at item
  if len(cmd) == 1 {
    LookInRoom(cmd, session)
  } else {
    LookAtItem(cmd, session)
  }
}

func NewLookCommand() CommandType {
  c := NewCommand("look", "'look' - Get information on current room.")
  c.Handler = LookCommandHandler
  AllCommandsBig["look"] =
  "Usage: 'look'\n" +
  "Gives you a detailed description of the area you're currently standing in.\n"

  CommandsHelp[len(CommandsHelp)] = "look"

  return c
}
