package Command

import (
  "fmt"

  "github.com/go-playground/validator/v10"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/models"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/arg"
)

type LoginCommand struct {
  Help string
}

type LoginInput struct {
  Username string `validate:"required,lte=255"`
  Password string `validate:"required,gte=8,lte=255"`
}

func NewLoginCommand() CommandType {
  c := NewCommand("login", "'login <username> <password>' - Log in as a character.'")
  c.Handler = LoginCommandHandler
  AllCommandsBig["login"] =
  "Usage: 'login <username> <password>'\n" +
  "Attempts to log into the <username> account with the <password> credential.\n" +
  "Example: 'login dock NotMyActualPassword'\n"

  CommandsHelp[len(CommandsHelp)] = "login"

  return c
}

func LoginCommandHandler(cmd []string, session *Game.Session) {
  if session.Authed {
    Arg.WriteFull(session.Conn, "You can't do this while you're logged in.\n")
    return
  }

  // Validate length of command
  if len(cmd) != 3 {
    fmt.Println("[LOGIN FAILURE (INVALID COMMAND)]", session.IP)
    Arg.Write(session.Conn, "Error! Be sure to use the following syntax...\n")
    Arg.WriteFull(session.Conn, "login <username> <password>\n")
    return
  }

  // Validate arguments in command
  validate := validator.New()

  input := &LoginInput{
    Username: cmd[1],
    Password: cmd[2],
  }

  err := validate.Struct(input)

  if err != nil {
    fmt.Printf("[LOGIN FAILURE (VALIDATION ERROR FOR '%s')] %s\n", cmd[1], session.IP)
    fmt.Println("[LOGIN FAILURE (VALIDATION ERROR)]", session.IP)
    Arg.Write(session.Conn, "Error! Be sure to use the following syntax...\n")
    Arg.WriteFull(session.Conn, "login <username> <password>\n")
    return
  }

  // Check if user exists
  searchUser := Model.User{}
  result := Data.DB.Where(Model.User{Username: cmd[1]}).First(&searchUser)

  if result.RowsAffected == 0 {
    // Let user and server know login was successful
    Arg.WriteFull(session.Conn, "I'm sorry, no user exists with that username.\n")

    fmt.Printf("[LOGIN FAILURE (USERNAME '%s' NOT FOUND)] %s\n", cmd[1], session.IP)

    return
  }

  // Validate Password
  valid := Model.ValidatePassword(cmd[2], searchUser.Password)

  if !valid {
    Arg.WriteFull(session.Conn, "That password is incorrect.\n")

    fmt.Printf("[LOGIN FAILURE (INCORRECT PASSWORD FOR '%s')] %s\n", cmd[1], session.IP)

    return
  }

  // Let user and server know login was successful
  Arg.WriteFull(session.Conn, "You've logged in! Welcome to Tenland. Begin by typing 'look'...\n")

  session.Authed = true
  session.User = searchUser
  Game.Sessions[session.ID] = *session // update sessions in list

  fmt.Printf("[LOGIN SUCCESS (FOR '%s')] %s\n", cmd[1], session.IP)
}
