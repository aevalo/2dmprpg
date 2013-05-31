package main

import (
	"fmt"
	"net"
	"2dmprpg/protocol"
)

func main() {
	// connect tcp
	fmt.Println("Opening a connection to 0.0.0.0:8000...")
	conn, err := net.Dial("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("Error occured while connecting!:", err)
		return
	}
	if conn != nil {
		// write data
		fmt.Println("Writing data...")
		_, err := protocol.WriteCommands(&conn, protocol.NewCommand("SESS", "MP"), protocol.NewCommand("LANG", "<>"))
		if err != nil {
			fmt.Println("Failed to send data:", err)
		}

		// read data
		fmt.Println("Reading data...")
		cmd := protocol.ReadCommand(&conn)
		fmt.Printf("Command: Name: %s, Data: %s\n", cmd.Name, cmd.Data)

		// close sockets
		fmt.Println("Closing connection...")
		err = conn.Close()
		if err != nil {
			fmt.Println("Failed to close connection:", err)
		}
	}
}
