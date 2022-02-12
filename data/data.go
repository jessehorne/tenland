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
Colors.Yellow(`Tenland is a multiplayer text-based RPG.
more coming soon...

Type 'help register' and 'help login' to begin.`)

var Begin =
Colors.Yellow("To begin, you need to register an account with us.\n" +
"register <username> <password>\n")

var Abyss =
Colors.Yellow("There is nothing around you except for a seemingly infinite stretch of grey.\n" +
"You can see your body and you can walk along the platform. How did you get here?\n")

var InvalidMovement =
Colors.Red("You can't go that way!\n")
