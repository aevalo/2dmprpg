package main

import (
	"fmt"
	"net"
	"os"
	"2dmprpg/server"
	"strconv"
	"strings"
)

func HandleConnection(conn *net.TCPConn, ch chan *server.Command) {
	quit := true
	for quit {
		fmt.Println(server.ReadCommand(conn).String())
	}
	conn.Close()
}

func Write(input string, conn *net.TCPConn) {
	_, err := server.WriteCommands(conn, server.NewCommand(input[:4], input[4:]))
	if err != nil {
		fmt.Println("Failed to send data:", err)
	}
}

func main() {
	args := os.Args[1:]
	ip := "localhost"
	port := 8000
	if len(args) > 2 {
		ip = args[0]
		port, _ = strconv.Atoi(args[1])
	}
	// connect tcp
	fmt.Println("Opening a connection ", ip, port)
	conn, err := net.DialTCP("tcp", nil, server.NewTCPAddr(ip, port))
	if err != nil {
		fmt.Println("Error occured while connecting!:", err)
		return
	}
	if conn != nil {

		ch := make(chan *server.Command)

		go HandleConnection(conn, ch)

		var command string
		var data string

		for {
			command = ""
			data = ""
			fmt.Print(">:")
			_, err := fmt.Scanf("%s%s", &command, &data)
			if err != nil {
				fmt.Println("Failed to read Stdin: ", err)
			} else {
				fmt.Println(command, data)
				server.WriteCommands(conn, server.NewCommand(strings.TrimSpace(command), strings.TrimSpace(data)))
			}
		}

		// close sockets
		fmt.Println("Closing connection...")
		err = conn.Close()
		if err != nil {
			fmt.Println("Failed to close connection:", err)
		}
	}
}
