package main

import (
  "fmt"
  "net"
  "sync"
  "2dmprpg/protocol"
)

func main() {
  fmt.Println("Opening a connection to 0.0.0.0:8000...")
  conn, err := net.Dial("tcp", "0.0.0.0:8000")
  if err != nil {
    fmt.Println("Error occured while connecting!:", err)
    return
  }
  if conn != nil {
    var wg sync.WaitGroup
    // Increment the WaitGroup counter
    //wg.Add(1)
    //go func(conn *net.Conn) {
      // Start session negotiation...
      fmt.Println("Writing data...")
      wg.Add(1)
      go func(conn *net.Conn) {
        _, err := protocol.WriteCommands(conn, protocol.NewCommand("SESS", "MP"), protocol.NewCommand("LANG", "<>"))
        if err != nil {
          fmt.Println("Failed to send data:", err)
        }
        wg.Done()
      }(&conn)
      wg.Wait()

      fmt.Println("Reading data...")
      wg.Add(1)
      go func(conn *net.Conn) {
        cmds := protocol.ReadCommands(conn)
        for i := range cmds {
          fmt.Printf("Command #%d: Name: %s, Data: %s\n", i, cmds[i].Name, cmds[i].Data)
        }
        wg.Done()
      }(&conn)
      wg.Wait()
      //n, err = protocol.WriteCommands(conn, protocol.NewCommand("VERS", "<>"),
      //                                      protocol.NewCommand("EXT0", "off"))
      //if err != nil {
      //  fmt.Println("Failed to send data:", err)
      //}
      //if n != 2 {
      //  fmt.Println("Not all commands were sent!")
      //}
      // Decrement the WaitGroup counter
      //wg.Done()
    //}(&conn)
    // Wait for the handler to complete
    //wg.Wait()
    // Close connection.
    fmt.Println("Closing connection...")
    err = conn.Close()
    if err != nil {
      fmt.Println("Failed to close connection:", err)
    }
  }
}
