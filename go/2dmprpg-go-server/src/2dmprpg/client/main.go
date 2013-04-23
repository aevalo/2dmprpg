package main

import (
  "log"
  "net"
  "io"
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
    _, err := io.WriteString(conn, "Hello, World!!!\n")
    if err != nil {
      log.Println("Failed to send data:", err)
    }
    log.Println("Closing connection...")
    err = conn.Close()
    if err != nil {
      log.Println("Failed to close connection:", err)
    }
  }
}

