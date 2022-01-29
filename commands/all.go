package Command

var Run = map[string]interface{}{
  "help": NewHelpCommand(),
  "register": NewRegisterCommand(),
  "login": NewLoginCommand(),
  "coords": NewCoordsCommand(),
  "look": NewLookCommand(),
}

type CommandType struct {
  Key string
  Help string
  Handler interface{}
}

func NewCommand(key string, help string) CommandType {
  nc := CommandType{}
  nc.Help = help

  return nc
}
