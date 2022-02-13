package Command

var AllCommands = make(map[string]string)
var AllCommandsBig = make(map[string]string)

var Run = map[string]interface{}{
  "help": NewHelpCommand(),
  "register": NewRegisterCommand(),
  "login": NewLoginCommand(),
  "coords": NewCoordsCommand(),
  "look": NewLookCommand(),
  "commands": NewCommandsCommand(),
  "begin": NewBeginCommand(),
  "build": NewBuildCommand(),
  "north": NewNorthCommand(),
  "south": NewSouthCommand(),
  "east": NewEastCommand(),
  "west": NewWestCommand(),
  "global": NewGlobalCommand(),
  "say": NewSayCommand(),
  "worth": NewWorthCommand(),
  "players": NewPlayersCommand(),
  "market": NewMarketCommand(),
  "item": NewItemCommand(),
  "inventory": NewInventoryCommand(),
  "me": NewMeCommand(),
  "take": NewTakeCommand(),
  "drop": NewDropCommand(),
}

type CommandType struct {
  Key string
  Help string
  Handler interface{}
}

func GetClosestMatch(search string) string {
  // loop through all full-length commands
  for key, _ := range Run {
    // Check if length of search is smaller than or equal to length of key
    // We do this because if a user types something longer than the command,
    // automatically know that it isn't it.
    if len(search) > len(key) {
      continue
    }

    // Check if search is equal to key starting at 0 and going to len(search)-1
    if search == key[:len(search)] {
      return key
    }
  }

  // Nothing found, return empty string
  return ""
}

func NewCommand(key string, help string) CommandType {
  nc := CommandType{}

  // Add to list of commands
  AllCommands[key] = help

  return nc
}
