package Command

import (
  "net"
  "fmt"

  "github.com/go-playground/validator/v10"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/models"
)

type RegisterCommand struct {
  Help string
}

type RegisterUserInput struct {
  Username string `validate:"required,lte=255"`
  Password string `validate:"required,gte=8,lte=255"`
}

func NewRegisterCommand() CommandType {
  rc := NewCommand("register", "register <username> <password>")
  rc.Handler = RegisterCommandHandler

  return rc
}

func RegisterCommandHandler(cmd []string, conn net.Conn) {
  // Validate length of command
  if len(cmd) != 3 {
    fmt.Println("[REGISTER FAILURE (INVALID COMMAND)]", conn.LocalAddr().String())
    conn.Write([]byte("Error! Use the following syntax...\n"))
    conn.Write([]byte("/register <username> <password>\n"))
    conn.Write([]byte(Data.Cursor))
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
    fmt.Println("[REGISTER FAILURE (INVALID SYNTAX)]", conn.LocalAddr().String())
    conn.Write([]byte("Error! Use the following syntax...\n"))
    conn.Write([]byte("/register <username> <password>\n"))
    conn.Write([]byte(Data.Cursor))
    return
  }

  // Create user
  Model.NewUser(cmd[1], cmd[2])


  // Let user and server know registration was successful
  conn.Write([]byte("You've registered. Now do 'login <username> <password>'!\n"))
  conn.Write([]byte(Data.Cursor))

  fmt.Println("[REGISTER SUCCESS]", conn.LocalAddr().String())
}
