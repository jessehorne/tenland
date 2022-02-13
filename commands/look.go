package Command

import (
  "fmt"
  "strconv"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/colors"
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
    session.Conn.Write([]byte("There is no such item in the abyss.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }


  // Get list of items in room
  allItems := Model.ItemsGetByRoom(Data.DB, x, y)

  for _,v := range allItems {
    if v.Name == cmd[1] {
      session.Conn.Write([]byte(fmt.Sprintf("You look closely at %s (%.2fkg)...\n", v.Name, v.Weight)))
      session.Conn.Write([]byte(v.Description + "\n"))
      session.Conn.Write([]byte(Data.Cursor))

      return
    }
  }

  session.Conn.Write([]byte("There doesn't appear to be anything like that here.\n"))
  session.Conn.Write([]byte(Data.Cursor))

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
  session.Conn.Write([]byte(searchRoom.Title + "\n"))
  session.Conn.Write([]byte(searchRoom.Desc + "\n\n"))
  session.Conn.Write([]byte(Colors.Yellow("Users: " + users + "\n")))
  session.Conn.Write([]byte("Items: " + items + "\n"))
  session.Conn.Write([]byte("Exits: [" + searchRoom.Exits + "]\n\n"))

  session.Conn.Write([]byte(Data.Cursor))
}

func LookCommandHandler(cmd []string, session *Game.Session) {
  // Make sure user is logged in
  if !session.Authed {
    session.Conn.Write([]byte("You see nothing. You have to be logged in to look around. Type 'help login' or 'help register'.\n"))
    session.Conn.Write([]byte(Data.Cursor))

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
  c := NewCommand("look", "'look' - Get visual and auditory information on current room.")
  c.Handler = LookCommandHandler
  AllCommandsBig["look"] =
  "Usage: 'look'\n" +
  "Gives you a detailed description of the area you're current standing in.\n"

  return c
}
