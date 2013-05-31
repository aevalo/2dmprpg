package server

import (
	"net"
	"strconv"
	"bytes"
	"fmt"
	"log"
	//	"bufio"
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

/*
 Reads command from tcp buffer. Blocking function.
*/
func ReadCommand(conn *net.TCPConn) *Command {
	//	reader := bufio.NewReader(*conn)
	buf := make([]byte, 4)
	var err error = nil
	_, err = conn.Read(buf)
	if err == nil {
		log.Printf("Read length %v\n", string(buf))
		data_len, err := strconv.Atoi(string(bytes.TrimSpace(buf)))
		_, err = conn.Read(buf)
		cmd := NewCommand(string(buf), "")
		log.Printf("Read commnd %v\n", string(buf))
		if err == nil {
			data_buf := make([]byte, data_len)
			_, err = conn.Read(data_buf)
			log.Printf("Read data %v\n", string(data_buf))
			cmd.Data = string(data_buf)
			return cmd
		} else {
			log.Printf("Error occurred while parsing data: %v\n", err)
		}
	}
	return nil
}

/*
 Writes one or more commands to tcp buffer. Non-blocking.
*/
func WriteCommands(conn *net.TCPConn, cmds ...*Command) (int, error) {
	//	conn := bufio.NewConn(*conn)
	var all int = 0
	for i := range cmds {
		n, err := conn.Write(cmds[i].Bytes())
		all = all + n
		if err != nil {
			//			_ = conn.Flush()
			return all, err
		}
		//		err = conn.Flush()
		if err != nil {
			return all, err
		}
	}
	return all, nil
}
