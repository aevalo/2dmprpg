package protocol

import (
  "net"
  "log"
  "strconv"
  "bytes"
  "fmt"
)


// Package initialization function
func init() {
  // Currently empty
}

type Command struct {
  Name string
  Data string
}

func NewCommand(name string, data string) *Command {
  if len(name) == 0 {
    return nil
  }
  c := new(Command)
  c.Name = name
  c.Data = data
  return c
}

func (c *Command) String() string {
  return fmt.Sprintf("%s%4d%s", c.Name, len(c.Data), c.Data)
}

func (c *Command) Bytes() []byte {
  buf := bytes.NewBufferString(c.String())
  return buf.Bytes()
}

func ParseCommand(conn net.Conn) *Command {
  buf := make([]byte, 4)

  n, err := conn.Read(buf)
  if err != nil || n != 4 {
    log.Println("Failed to read command:", err)
    return nil
  }
  cmd := string(buf)

  n, err = conn.Read(buf)
  if err != nil || n != 4 {
    log.Println("Failed to read data length:", err)
    return nil
  }
  data_len, err := strconv.Atoi(string(bytes.TrimSpace(buf)))
  if err != nil {
    log.Println("Failed to convert data length:", err)
    return nil
  }

  data_buf := make([]byte, data_len)
  n, err = conn.Read(data_buf)
  if err != nil || n != data_len {
    log.Println("Failed to read data:", err)
    return nil
  }
  data := string(data_buf)

  return &Command{Name: cmd, Data: data}
}

