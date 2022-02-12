package Command

import (
  "strconv"
  "fmt"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

func CreateCommandHandler(cmd []string, session *Game.Session) {
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

  // Validate that cmd is at least 5 in length
  if len(cmd) < 5 {
    session.Conn.Write([]byte("Invalid. Incorrect number of arguments. Please type 'help create'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }


  // Get arguments and set appropriate variables
  name := cmd[1]
  x, _ := strconv.Atoi(cmd[2])
  y, _ := strconv.Atoi(cmd[3])
  held := false
  userID := -1

  if cmd[4] == "true" {
    // Validate that cmd is at least 6 in length
    if len(cmd) < 6 {
      session.Conn.Write([]byte("Invalid. You have to specify a User ID. Please type 'help create'.\n"))
      session.Conn.Write([]byte(Data.Cursor))
      return
    }

    held = true
    userID, _ = strconv.Atoi(cmd[5])
  } else {
    held = false
  }

  // Create item
  if held {
    newItem, err := Model.ItemNew(Data.DB, name, x, y, held, uint(userID))

    if err != nil {
      session.Conn.Write([]byte("There was an error..."))
      session.Conn.Write([]byte(err.Error()))
      session.Conn.Write([]byte("\n" + Data.Cursor))
      return
    }

    session.Conn.Write([]byte(fmt.Sprintf("Created item named '%s'!", newItem.Name)))
    session.Conn.Write([]byte("\n" + Data.Cursor))
  } else {
    // If not creating item in inventory, create it on the ground but set user
    // id to the current logged in user so can know who created the item
    newItem, err := Model.ItemNew(Data.DB, name, x, y, held, session.User.ID)

    if err != nil {
      session.Conn.Write([]byte("There was an error..."))
      session.Conn.Write([]byte(err.Error()))
      session.Conn.Write([]byte("\n" + Data.Cursor))
      return
    }

    session.Conn.Write([]byte(fmt.Sprintf("Created item named '%s'!", newItem.Name)))
    session.Conn.Write([]byte("\n" + Data.Cursor))
  }

  // end of command
}

func NewCreateCommand() CommandType {
  hc := NewCommand("create", "'create <name> <x> <y> <held> <userid?>' - Creates an item.")
  hc.Handler = CreateCommandHandler
  AllCommandsBig["create"] = `
Usage: 'create <name> <x> <y> <held> <userid?>
Creates an item named <name> with an origin position of <x>,<y>. If <held> is
true, <userid> is necessary as well.

Example: create SuperSword 0 0 false
That will create an item named 'SuperSword' and place it in the world at 0,0.

Example 2: create ConquestOfBreadBook 0 0 true 1
That will create an item named 'ConquestOfBreadBook' and give it to the player
whose ID is '1'.
  `

  return hc
}
