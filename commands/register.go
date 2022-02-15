package Command

import (
  "fmt"

  "github.com/go-playground/validator/v10"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/arg"
)

type RegisterCommand struct {
  Help string
}

type RegisterUserInput struct {
  Username string `validate:"required,lte=255"`
  Password string `validate:"required,gte=8,lte=255"`
}

func NewRegisterCommand() CommandType {
  rc := NewCommand("register", "'register <username> <password>' - Register new character.'")
  rc.Handler = RegisterCommandHandler
  AllCommandsBig["register"] =
  "Usage: 'register <username> <password>'\n" +
  "Attempts to register the <username> account with the <password> credential.\n" +
  "Example: 'register dock NotMyActualPassword'\n"

  CommandsHelp[len(CommandsHelp)] = "register"

  return rc
}

func RegisterCommandHandler(cmd []string, session *Game.Session) {
  if session.Authed {
    Arg.WriteFull(session.Conn, "You can't do this while you're logged in.\n")
    return
  }

  // Validate length of command
  if len(cmd) != 3 {
    fmt.Println("[REGISTER FAILURE (INVALID COMMAND)]", session.IP)
    Arg.Write(session.Conn, "Error! Use the following syntax...\n")
    Arg.WriteFull(session.Conn, "register <username> <password>\n")
    return
  }

  // Validate arguments in command
  validate := validator.New()

  input := &RegisterUserInput{
    Username: cmd[1],
    Password: cmd[2],
  }

  err := validate.Struct(input)

  if err != nil {
    fmt.Println("[REGISTER FAILURE (VALIDATION ERROR)]", session.IP)
    Arg.Write(session.Conn, "Error! Use the following syntax...\n")
    Arg.Write(session.Conn, "/register <username> <password>\n")
    Arg.WriteFull(session.Conn, "Your username must be unique and your password must be at least 8 characters.\n")
    return
  }

  // Create user
  _, userCreationError := Model.NewUser(Data.DB, cmd[1], cmd[2])

  if userCreationError != nil {
    fmt.Println("[REGISTER FAILURE (CREATION ERROR)]", session.IP)

    Arg.Write(session.Conn, "I'm sorry, something went wrong.\n")
    Arg.WriteFull(session.Conn, userCreationError.Error())

    return
  }

  // Let user and server know registration was successful
  Arg.WriteFull(session.Conn, "You've registered. Now do 'login <username> <password>'!\n")

  fmt.Println("[REGISTER SUCCESS]", session.IP)
}
