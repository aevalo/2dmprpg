package protocol

import (
  "net"
  "log"
  "strconv"
  "bytes"
  "fmt"
  "bufio"
  "io/ioutil" // For ReadAll 
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
  arr := make([]*Command, 0)
  allbytes, err := ioutil.ReadAll(bufio.NewReader(conn))
  if err != nil {
    log.Println("Failed to read commands:", err)
    return arr
  }
  buffer := bytes.NewBuffer(allbytes)

  for buffer.Len() > 0 {
    buf := make([]byte, 4)

    n, err := buffer.Read(buf)
    if err != nil || n != 4 {
      log.Println("Failed to read command:", err)
      return nil
    }
    cmd := string(buf)

    n, err = buffer.Read(buf)
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
    n, err = buffer.Read(data_buf)
    if err != nil || n != data_len {
      log.Println("Failed to read data:", err)
      return nil
    }
    data := string(data_buf)

    arr = append(arr, &Command{Name: cmd, Data: data})
  }
  return arr
}

func WriteCommandsArray(conn net.Conn, cmds []*Command) (int, error) {
  var sent int = 0
  for i := range cmds {
    log.Println("Writing command", cmds[i].Name)
    _, err := conn.Write(cmds[i].Bytes())
    log.Println(err)
    if err == nil {
      sent++
    } else {
      if err != nil {
        log.Println("Failed to send data:", err)
      }
      return sent, err
    }
  }
  log.Println("All went well!")
  return sent, nil
}

func WriteCommands(conn net.Conn, cmds ...*Command) (int, error) {
  var sent int = 0
  for i := range cmds {
    bytes := cmds[i].Bytes()
    log.Println(bytes)
    n, err := conn.Write(bytes)
    if err == nil && n == len(bytes) {
      sent++
    } else {
      if err != nil {
        log.Println("Failed to send data:", err)
      }
      if n != len(bytes) {
        log.Println("Not all data was sent!")
      }
      return sent, err
    }
  }
  return sent, nil
}