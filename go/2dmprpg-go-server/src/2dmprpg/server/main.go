package main

import (
  "log"
  "net"
  "sync"
  "time"
  "2dmprpg/protocol"
)

func handleConnection(conn net.Conn) {
  cmds := protocol.ReadCommands(conn)
  for i := range cmds {
    log.Printf("Command #%d: Name: %s, Data: %s\n", i, cmds[i].Name, cmds[i].Data)
  }
  log.Println("Writing data...")
  _, err := protocol.WriteCommands(conn, protocol.NewCommand("SESS", "MP"), protocol.NewCommand("LANG", "<>"))
  if err != nil {
    log.Println("Failed to send data:", err)
  }

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
  log.Println("Listening 0.0.0.0:8000 for connections...")
  ln, err := net.Listen("tcp", "0.0.0.0:8000")
  if err != nil {
    log.Println("Error occured while listening for connections:", err)
    return
  }
  if ln != nil {
    log.Println("Waiting for incoming connections...")
    conn, err := ln.Accept()
    if err != nil {
      log.Println("Failed to accept connection:", err)
      return
    } else {
      if conn != nil {
        var wg sync.WaitGroup

        log.Println("Handling connection...")
        // Increment the WaitGroup counter
        wg.Add(1)

        // Start the coroutine to handle connection
        go func(conn net.Conn) {
          handleConnection(conn)
          err = conn.Close()
          if err != nil {
            log.Println("Failed to close connection:", err)
          }
          // Decrement the WaitGroup counter
          wg.Done()
        }(conn)

        // Wait for the handler to complete
        wg.Wait()

        time.Sleep(time.Millisecond * 1000)

        log.Println("Done!")
      }
    }
  }
}
