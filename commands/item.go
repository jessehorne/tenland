package Command

import (
  "strconv"
  "fmt"
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/arg"
)

func WeightItem(cmd []string, session *Game.Session) {
  if len(cmd) < 4 {
    Arg.WriteFull(session.Conn, "Invalid. Incorrect number of arguments. Please type 'help item'.\n")
    return
  }

  itemID,_ := strconv.Atoi(cmd[2])
  weight, err := strconv.ParseFloat(cmd[3], 32)

  if err != nil {
    Arg.WriteFull(session.Conn, "Invalid. Incorrect weight value. Please type 'help item'.\n")
    return
  }

  success := Model.ItemUpdateWeight(Data.DB, uint(itemID), float32(weight))

  if !success {
    Arg.WriteFull(session.Conn, "Invalid. Could not update item weight in the database. Please type 'help item'.\n")
    return
  }

  Arg.WriteFull(session.Conn, "You've updated an items weight!\n")
}

func DescribeItem(cmd []string, session *Game.Session) {
  if len(cmd) < 4 {
    Arg.WriteFull(session.Conn, "Invalid. Incorrect number of arguments. Please type 'help item'.\n")
    return
  }

  itemID,_ := strconv.Atoi(cmd[2])
  description := strings.Join(cmd[3:], " ")

  success := Model.ItemUpdateDescription(Data.DB, uint(itemID), description)

  if !success {
    Arg.WriteFull(session.Conn, "Invalid. Could not update item description in the database. Please type 'help item'.\n")
    return
  }

  Arg.WriteFull(session.Conn, "You've updated an items description!\n")
}

func CreateItem(cmd []string, session *Game.Session) {
  if len(cmd) < 6 {
    Arg.WriteFull(session.Conn, "Invalid. Incorrect number of arguments. Please type 'help item'.\n")
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
      Arg.WriteFull(session.Conn, "Invalid. You have to specify a User ID. Please type 'help item'.\n")
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
      Arg.Write(session.Conn, "There was an error...\n")
      Arg.WriteFull(session.Conn, err.Error())
      return
    }

    Arg.WriteFull(session.Conn, fmt.Sprintf("Created item named '%s'!\n", newItem.Name))
  } else {
    // If not creating item in inventory, create it on the ground but set user
    // id to the current logged in user so can know who created the item
    newItem, err := Model.ItemNew(Data.DB, name, x, y, held, session.User.ID)

    if err != nil {
      Arg.Write(session.Conn, "There was an error...\n")
      Arg.WriteFull(session.Conn, err.Error())
      return
    }

    Arg.WriteFull(session.Conn, fmt.Sprintf("Created item named '%s'!\n", newItem.Name))
  }
}

func ItemCommandHandler(cmd []string, session *Game.Session) {
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

    // Verify that there is at least one arg to specify the command
    if len(cmd) < 2 {
      Arg.WriteFull(session.Conn, "You need to supply a command to 'item'. Please type 'help item'.\n")
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
      Arg.WriteFull(session.Conn, "Invalid command supplied to 'item'. Please type 'help item'.\n")
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
