package Command

var AllCommands = make(map[string]string)

var Run = map[string]interface{}{
  "help": NewHelpCommand(),
  "register": NewRegisterCommand(),
  "login": NewLoginCommand(),
  "coords": NewCoordsCommand(),
  "look": NewLookCommand(),
  "commands": NewCommandsCommand(),
}

type CommandType struct {
  Key string
  Help string
  Handler interface{}
}

func NewCommand(key string, help string) CommandType {
  nc := CommandType{}
  nc.Help = help

  // Add to list of commands
  AllCommands[key] = help

  return nc
}
