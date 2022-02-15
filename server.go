package main

import (
  "net"
  "fmt"
  "os"
  "strings"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/arg"
  "github.com/jessehorne/tenland/game"
  "github.com/jessehorne/tenland/commands"

  "github.com/joho/godotenv"
)

func main() {
  // Load environment variables
  err := godotenv.Load()

  if err != nil {
    fmt.Println("[ERROR] Error loading .env file...")
  }

  // Setup MySQL Database

  Data.InitDB()

  fmt.Println("Database initialized...")

  // Set up server
  service := ":4000"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)

  checkError(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)

  fmt.Println("Tenland v0.0.1 running on port 4000")

  // Setup ticker
  Game.StartTicker()

  for {
    conn, err := listener.Accept()

    if err != nil {
      continue
    }

    go handleClient(conn)
  }
}

func handleClient(conn net.Conn) {
  defer conn.Close()

  var buf [512]byte

  // Tell server user connected
  fmt.Println("[USER CONNECTED]", conn.LocalAddr().String())

  // Give user welcome
  Arg.WriteFull(conn, Data.Welcome + "\n")

  // Create session
  session := Game.NewSession(conn)

  for {
    // Read input
    n, err := conn.Read(buf[0:])

    if err != nil {
      return
    }

    // Compare input
    Handle(n, buf, &session)

  }
}

func Handle(n int, buf [512]byte, session *Game.Session) {
  // Split command
  cmd := string(buf[0:n-1])
  splitCmd := strings.Split(cmd, " ")

  if cmd == "" {
    Arg.Cursor(session.Conn)
    return
  }

  if splitCmd[0] == "exit" {
    Arg.Write(session.Conn, Data.Goodbye)
    session.Conn.Close()
    fmt.Println("[USER DISCONNECTED]", session.IP)
  } else {
    // Get closest match
    match := Command.GetClosestMatch(splitCmd[0])

    // No match found
    if match == "" {
      Arg.WriteFull(session.Conn, Data.UnknownCommand)
      return
    }

    // Get command from Run map
    f, found := Command.Run[match]

    if found {
      f.(Command.CommandType).Handler.(func([]string, *Game.Session))(splitCmd, session)
    } else {
      Arg.WriteFull(session.Conn, Data.UnknownCommand)
    }
  }
}

func checkError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())

    os.Exit(1)
  }
}
