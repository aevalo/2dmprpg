package main

import (
  "fmt"
  "net"
  "sync"
  "2dmprpg/protocol"
)

func handleConnection(conn net.Conn) {
  cmds := protocol.ReadCommands(conn)
  for i := range cmds {
    fmt.Printf("Command #%d: Name: %s, Data: %s\n", i, cmds[i].Name, cmds[i].Data)
  }
  fmt.Println("Writing data...")
  arr := []*protocol.Command{protocol.NewCommand("SESS", "MP"), protocol.NewCommand("LANG", "<>")}
  _, _ = protocol.WriteCommandsArray(conn, arr)

  //n, err := protocol.WriteCommands(conn,  protocol.NewCommand("VERS", "<>"),
  //                                        protocol.NewCommand("LANG", "<>"),
   //                                       protocol.NewCommand("EXT0", "test"),
   //                                       protocol.NewCommand("PALT", "file://tmp/mypalette.pal"))
  //if err != nil {
  //  fmt.Println("Failed to send data:", err)
  //}
  //if n != 4 {
  //  fmt.Println("Not all commands were sent!")
  //}
  //cmds = protocol.ReadCommands(conn)
  //for i := range cmds {
  //  fmt.Printf("Command #%d: Name: %s, Data: %s\n", i, cmds[i].Name, cmds[i].Data)
  //}
}

func main() {
  fmt.Println("Listening 127.0.0.1:8000 for connections...")
  ln, err := net.Listen("tcp", "127.0.0.1:8000")
  if err != nil {
    fmt.Println("Error occured while listening for connections:", err)
    return
  }
  if ln != nil {
    fmt.Println("Waiting for incoming connections...")
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Failed to accept connection:", err)
      return
    } else {
      if conn != nil {
        var wg sync.WaitGroup

        fmt.Println("Handling connection...")
        // Increment the WaitGroup counter
        wg.Add(1)

        // Start the coroutine to handle connection
        go func(conn net.Conn) {
          handleConnection(conn)
          err = conn.Close()
          if err != nil {
            fmt.Println("Failed to close connection:", err)
          }
          // Decrement the WaitGroup counter
          wg.Done()
        }(conn)

        // Wait for the handler to complete
        wg.Wait()

        fmt.Println("Done!")
      }
    }
  }
}
