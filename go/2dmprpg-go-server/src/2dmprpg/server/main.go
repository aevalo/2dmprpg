package main

import (
  "log"
  "net"
  "sync"
  "2dmprpg/protocol"
)

func handleConnection(conn *net.Conn) {
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

        log.Println("Handling connection...")
        var wg sync.WaitGroup
        // Increment the WaitGroup counter
        //wg.Add(1)

        // Start the coroutine to handle connection
        //go func(conn *net.Conn) {
        //  handleConnection(conn)
          // Decrement the WaitGroup counter
        //  wg.Done()
        //}(&conn)

        log.Println("Reading data...")
        wg.Add(1)
        go func(conn *net.Conn) {
          cmds := protocol.ReadCommands(conn)
          for i := range cmds {
            log.Printf("Command #%d: Name: %s, Data: %s\n", i, cmds[i].Name, cmds[i].Data)
          }
          wg.Done()
        }(&conn)
        wg.Wait()

        log.Println("Writing data...")
        wg.Add(1)
        go func(conn *net.Conn) {
          _, err := protocol.WriteCommands(conn, protocol.NewCommand("SESS", "MP"), protocol.NewCommand("LANG", "<>"))
          if err != nil {
            log.Println("Failed to send data:", err)
          }
          wg.Done()
        }(&conn)
        wg.Wait()


        log.Println("Closing connection...")
        err = conn.Close()
        if err != nil {
          log.Println("Failed to close connection:", err)
        }

        log.Println("Done!")
      }
    }
  }
}
