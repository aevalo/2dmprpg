package main

import (
  "log"
  "net"
  "sync"
  "2dmprpg/protocol"
)

func handleConnection(conn net.Conn) {
  cmd := protocol.ParseCommand(conn)
  log.Println("Command:", cmd.Name)
  log.Println("Data:", cmd.Data)
  cmd = protocol.ParseCommand(conn)
  log.Println("Command:", cmd.Name)
  log.Println("Data:", cmd.Data)
}

func main() {
  log.Println("Listening 127.0.0.1:8000 for connections...")
  ln, err := net.Listen("tcp", "127.0.0.1:8000")
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

        log.Println("Done!")
      }
    }
  }
}
