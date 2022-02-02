package Command

import (
  "fmt"

  "github.com/go-playground/validator/v10"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/game"
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
  rc.Help =
  "register <username> <password>\n" +
  "Attempts to register the <username> account with the <password> credential.\n" +
  "Example: 'register dock NotMyActualPassword'\n"

  return rc
}

func RegisterCommandHandler(cmd []string, session *Game.Session) {
  if session.Authed {
    session.Conn.Write([]byte("You can't do this while you're logged in.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  // Validate length of command
  if len(cmd) != 3 {
    fmt.Println("[REGISTER FAILURE (INVALID COMMAND)]", session.IP)
    session.Conn.Write([]byte("Error! Use the following syntax...\n"))
    session.Conn.Write([]byte("register <username> <password>\n"))
    session.Conn.Write([]byte(Data.Cursor))
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
    session.Conn.Write([]byte("Error! Use the following syntax...\n"))
    session.Conn.Write([]byte("/register <username> <password>\n"))
    session.Conn.Write([]byte("Your username must be unique and your password must be at least 8 characters.\n"))
    session.Conn.Write([]byte(Data.Cursor))
    return
  }

  // Create user
  _, userCreationError := Model.NewUser(Data.DB, cmd[1], cmd[2])

  if userCreationError != nil {
    fmt.Println("[REGISTER FAILURE (CREATION ERROR)]", session.IP)

    session.Conn.Write([]byte("I'm sorry, something went wrong.\n"))
    session.Conn.Write([]byte(userCreationError.Error()))
    session.Conn.Write([]byte(Data.Cursor))

    return
  }

  // Let user and server know registration was successful
  session.Conn.Write([]byte("You've registered. Now do 'login <username> <password>'!\n"))
  session.Conn.Write([]byte(Data.Cursor))

  fmt.Println("[REGISTER SUCCESS]", session.IP)
}
