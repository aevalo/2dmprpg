package protocol

import (
	"net"
	"strconv"
	"bytes"
	"fmt"
	"log"
	"bufio"
	"time"
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
	return fmt.Sprintf("%4d%s%s", len(c.Data), c.Name, c.Data)
}

func (c *Command) Bytes() []byte {
	buf := bytes.NewBufferString(c.String())
	return buf.Bytes()
}

func ReadCommands(conn *net.Conn) []*Command {
	cmds := make([]*Command, 0)
	reader := bufio.NewReader(*conn)
	buf := make([]byte, 4)
	var err error = nil
	for err == nil {
		tim := time.Time.Now()
		tim.Add(time.Duration(1) * time.Second)
		conn.SetDeadline(tim)
		log.Printf("start reading")
		_, err = reader.Read(buf)
		if err == nil {
			log.Printf("Read length %v\n", string(buf))
			data_len, err := strconv.Atoi(string(bytes.TrimSpace(buf)))
			_, err = reader.Read(buf)
			cmd := NewCommand(string(buf), "")
			log.Printf("Read commnd %v\n", string(buf))
			if err == nil {
				data_buf := make([]byte, data_len)
				_, err = reader.Read(data_buf)
				log.Printf("Read data %v\n", string(data_buf))
				cmd.Data = string(data_buf)
				cmds = append(cmds, cmd)
			} else {
				log.Printf("Error occurred while parsing data: %v\n", err)
			}
		}
	}
	return cmds
}

func WriteCommands(conn *net.Conn, cmds ...*Command) (int, error) {
	writer := bufio.NewWriter(*conn)
	var all int = 0
	for i := range cmds {
		log.Printf("start writing")
		n, err := writer.WriteString(cmds[i].String())
		all = all + n
		if err != nil {
			_ = writer.Flush()
			return all, err
		}
		err = writer.Flush()
		if err != nil {
			return all, err
		}
	}
	return all, nil
}
