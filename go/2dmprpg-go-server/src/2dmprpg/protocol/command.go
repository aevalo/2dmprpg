package protocol

import (
  "net"
  "strconv"
  "bytes"
  "strings"
  "fmt"
  "log"
  "bufio"
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

func ReadCommands(conn net.Conn) []*Command {
  cmds := make([]*Command, 0)
  reader := bufio.NewReader(conn)
  buf := make([]byte, 4)
  var n int = 0
  var err error = nil
  for err == nil {
    n, err = reader.Read(buf)
    if err == nil && n == 4 {
      cmd := NewCommand(string(buf), "")
      n, err = reader.Read(buf)
      data_len, err := strconv.Atoi(string(bytes.TrimSpace(buf)))
      if err == nil {
        data_buf := make([]byte, data_len)
        n, err = reader.Read(data_buf)
        cmd.Data = string(data_buf)
        cmds = append(cmds, cmd)
      }
    } else {
      log.Printf("Error occured while reading: %v\n", err)
    }
  }
  return cmds
}

func WriteCommandsArray(conn net.Conn, cmds []*Command) (int, error) {
  buf := make([]string, len(cmds))
  for i := range cmds {
    buf = append(buf, cmds[i].String())
  }
  return conn.Write([]byte(strings.Join(buf, "")))
}

func WriteCommands(conn net.Conn, cmds ...*Command) (int, error) {
  buf := make([]string, len(cmds))
  for i := range cmds {
    buf = append(buf, cmds[i].String())
  }
  return conn.Write([]byte(strings.Join(buf, "")))
}