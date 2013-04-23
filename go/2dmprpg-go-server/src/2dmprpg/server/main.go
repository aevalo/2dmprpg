package main

import (
  "log"
  "net"
  "bufio"
  "sync"
  "regexp"
)

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
  reader := bufio.NewReader(conn)
  str, err := reader.ReadString('\n')
  if err != nil {
    log.Println("Failed to read data:", err)
  } else {
    log.Println("Received data:", str)
    re, err := regexp.Compile("^\\s*(\\w+)\\s*(\\d+)(\\w+)\\s*$")
    if err != nil {
      log.Println("Error occured while compiling regular expression:", err)
    } else {
      matches := re.FindStringSubmatch(str)
      log.Println("Num of matches:", len(matches))
      for i := 0; i < len(matches); i++ {
        log.Println(i, ") ", matches[i])
      }
      log.Println(matches)
    }
  }
  err = conn.Close()
  if err != nil {
    log.Println("Failed to close connection:", err)
  }
  // Decrement the counter.
  wg.Done()
}

func main() {
  log.Println("Listening 127.0.0.1:8000 for connections...")
  ln, err := net.Listen("tcp", "127.0.0.1:8000")
  if err != nil {
    log.Println("Error occured while listening for connections:", err)
    return
  }
  if ln != nil {
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
        go handleConnection(conn, &wg)

        // Wait for the handler to complete
        wg.Wait()

        log.Println("Done!")
      }
    }
  }
}
