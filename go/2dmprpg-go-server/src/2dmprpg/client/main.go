package main

import (
  "log"
  "net"
  "io"
  "2dmprpg/protocol"
)

func main() {
  log.Println("Opening a connection to 127.0.0.1:8000...")
  conn, err := net.Dial("tcp", "127.0.0.1:8000")
  if err != nil {
    log.Println("Error occured while connecting!:", err)
    return
  }
  if conn != nil {
    log.Println("Connected, sending data...")
    cmd := protocol.NewCommand("SESS", "MP")
    n, err := io.WriteString(conn, cmd.String())
    if err != nil || n != len(cmd.String()) {
      log.Println("Failed to send data:", err)
    }
    n, err = conn.Write(cmd.Bytes())
    if err != nil || n != len(cmd.Bytes()) {
      log.Println("Failed to send data:", err)
    }
    log.Println("Closing connection...")
    err = conn.Close()
    if err != nil {
      log.Println("Failed to close connection:", err)
    }
  }
}
