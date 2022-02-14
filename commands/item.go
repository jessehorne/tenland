package Command

import (
  "strconv"
  "fmt"
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
)

func WeightItem(cmd []string, session *Game.Session) {
  if len(cmd) < 4 {
    session.Conn.Write([]byte("Invalid. Incorrect number of arguments. Please type 'help item'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  itemID,_ := strconv.Atoi(cmd[2])
  weight, err := strconv.ParseFloat(cmd[3], 32)

  if err != nil {
    session.Conn.Write([]byte("Invalid. Incorrect weight value. Please type 'help item'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  success := Model.ItemUpdateWeight(Data.DB, uint(itemID), float32(weight))

  if !success {
    session.Conn.Write([]byte("Invalid. Could not update item weight in the database. Please type 'help item'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  session.Conn.Write([]byte("You've updated an items weight!\n"))
  session.Conn.Write([]byte(Data.Cursor))
}

func DescribeItem(cmd []string, session *Game.Session) {
  if len(cmd) < 4 {
    session.Conn.Write([]byte("Invalid. Incorrect number of arguments. Please type 'help item'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  itemID,_ := strconv.Atoi(cmd[2])
  description := strings.Join(cmd[3:], " ")

  success := Model.ItemUpdateDescription(Data.DB, uint(itemID), description)

  if !success {
    session.Conn.Write([]byte("Invalid. Could not update item description in the database. Please type 'help item'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  session.Conn.Write([]byte("You've updated an items description!\n"))
  session.Conn.Write([]byte(Data.Cursor))
}

func CreateItem(cmd []string, session *Game.Session) {
  if len(cmd) < 6 {
    session.Conn.Write([]byte("Invalid. Incorrect number of arguments. Please type 'help item'.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }


  // Get arguments and set appropriate variables
  name := cmd[2]
  x, _ := strconv.Atoi(cmd[3])
  y, _ := strconv.Atoi(cmd[4])
  held := false
  userID := -1

  if cmd[4] == "true" {
    // Validate that cmd is at least 6 in length
    if len(cmd) < 7 {
      session.Conn.Write([]byte("Invalid. You have to specify a User ID. Please type 'help item'.\n"))
      session.Conn.Write([]byte(Data.Cursor))
      return
    }

    held = true
    userID, _ = strconv.Atoi(cmd[6])
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
}

func ItemCommandHandler(cmd []string, session *Game.Session) {
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

    // Verify that there is at least one arg to specify the command
    if len(cmd) < 2 {
      session.Conn.Write([]byte("You need to supply a command to 'item'. Please type 'help item'.\n"))
      session.Conn.Write([]byte(Data.Cursor))
      return
    }

    // Handle commands
    command := cmd[1]

    if command == "create" {
      CreateItem(cmd, session)
    } else if command == "describe" {
      DescribeItem(cmd, session)
    } else if command == "weight" {
      WeightItem(cmd, session)
    } else {
      session.Conn.Write([]byte("Invalid command supplied to 'item'. Please type 'help item'.\n"))
      session.Conn.Write([]byte(Data.Cursor))
    }

  // end of command
}

func NewItemCommand() CommandType {
  hc := NewCommand("item", "'item <command> <arg> ...' - Creates an item.")
  hc.Handler = ItemCommandHandler
  AllCommandsBig["item"] = `
Usage: 'item <cmd> <args> ...'

Options
=======

1. 'create' - Create an item. (See Creating an Item)
2. 'describe' - Describe an item. (See Describing an Item)

Create an Item
==============

Command: 'item create <name> <x> <y> <held> <userid?>'

Creates an item named <name> with an origin position of <x>,<y>. If <held> is
true, <userid> is necessary as well.

Example: item create SuperSword 0 0 false
That will create an item named 'SuperSword' and place it in the world at 0,0.

Example 2: item create ConquestOfBreadBook 0 0 true 1
That will create an item named 'ConquestOfBreadBook' and give it to the player
whose ID is '1'.

Describing an Item
==================

Command: 'item describe <itemID> <description>'

Sets the description of an item.

Example: 'item describe 1 A book written by P. Kropotkin explaining Anarchism.'

Give Weight to an Item
==================

Command: 'item weight <itemID> <float>'

Sets the weight of an item.

Example: 'item weight 1 2.5'

The above example sets the item with ID of '1' to 2.5kg.
`

  CommandsHelp[len(CommandsHelp)] = "item"

  return hc
}
