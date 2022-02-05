package Data

import (
  "github.com/jessehorne/tenland/colors"
)

var Cursor = Colors.Blue("tl>")

var Welcome =
Colors.Yellow("Welcome to Tenland (v0.0.1). To get started, type 'begin' or 'help'.\n\n")

var Goodbye =
Colors.Yellow("Thank you for playing Tenland. Come back soon! (press return)")

var UnknownCommand =
Colors.Red("I'm sorry, I do not understand what you mean. Try 'commands'.\n\n")

var Help =
Colors.Yellow(`Tenland is a multiplayer text-based game.
In Tenland, you start out as a ghost that can move around in the endless void
of empty space. To begin playing you must 'register' and then 'login'. Try
'help <command>' to get more information on commands. For example, type...
'help register'.
Once you're registered, you'll enter the world of Tenland. Tenland is a large
circular island. There are different resources on different parts of the island.
For a map of the island resources, type 'map'.
Tenland itself can be played in a variety of ways. Some players choose to gain
wealth in the form of Gold (type 'gold') while others choose to do PvP or build.
Please note that Tenland is very new. There's not a ton of information out yet.
Stay tuned and feel free to help by being active on the game, in the community
and by contributing code.
Try 'commands' for a list of all available commands.
Type 'begin' to start.`)

var Begin =
Colors.Yellow("To begin, you need to register an account with us.\n" +
"register <username> <password>\n")

var Abyss =
Colors.Yellow("There is nothing around you except for a seemingly infinite stretch of grey.\n" +
"You can see your body and you can walk along the platform. How did you get here?\n")

var InvalidMovement =
Colors.Red("You can't go that way!\n")
