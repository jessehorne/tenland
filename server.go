package main

import (
  "net"
  "fmt"
  "os"

  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/arg"
  "github.com/jessehorne/tenland/game"

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
  Arg.WriteFull([]byte(Data.Welcome), conn)

  // Create session
  session := Game.NewSession(conn)

  for {
    // Session debug
    fmt.Println("Session", session.Authed, session.IP)

    // Read input
    n, err := conn.Read(buf[0:])

    if err != nil {
      return
    }

    // Compare input
    Arg.Handle(n, buf, &session)

  }
}

func checkError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())

    os.Exit(1)
  }
}
