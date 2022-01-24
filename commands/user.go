package Command

import (
  "net"
  "fmt"

  "github.com/go-playground/validator/v10"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/models"
)

type RegisterUserInput struct {
  Username string `validate:"required,lte=255"`
  Password string `validate:"required,gte=8,lte=255"`
}

func RegisterCommandHandler(cmd []string, conn net.Conn) {
  // Validate length of command
  if len(cmd) != 3 {
    fmt.Println("[REGISTER FAILURE]", conn.LocalAddr().String())
    conn.Write([]byte("INCORRECT COMMAND!\n"))
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
    fmt.Println("[REGISTER FAILURE]", conn.LocalAddr().String())
    conn.Write([]byte("MISSING INPUT!\n"))
    conn.Write([]byte(Data.Cursor))
    return
  }

  // Create user
  newPlayer := Models.NewUser(cmd[1], cmd[2])


  // Let user and server know registration was successful
  conn.Write([]byte("YOU TRIED TO REGISTER SUCCESSFULLY!\n"))
  conn.Write([]byte("Username: "))
  conn.Write([]byte(newPlayer.Username))
  conn.Write([]byte("\n"))
  conn.Write([]byte(Data.Cursor))

  fmt.Println("[REGISTER SUCCESS]", conn.LocalAddr().String())
}
