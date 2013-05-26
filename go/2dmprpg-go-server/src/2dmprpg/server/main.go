package main

import (
	"log"
	"net"
	"2dmprpg/protocol"
	"time"
)

func handleConnection(conn *net.Conn, ch chan *protocol.Command) {
	//TODO: add some connection/user handling here

	// handle incoming command the best you can
	quit := false
	for !quit {
		cmd := protocol.ReadCommand(conn)
		log.Printf("Command: Name: %s, Data: %s\n", cmd.Name, cmd.Data)

		// TODO: add some quit / auth handling here

		ch <- cmd
	}
}

func main() {
	// start listening tcp connections
	log.Println("Listening 0.0.0.0:8000 for connections...")
	ln, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Println("Error occured while listening for connections:", err)
		return
	}
	if ln != nil {
		// waiting for connection
		log.Println("Waiting for incoming connections...")
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			return
		} else {
			if conn != nil {
				// Start message handling.
				log.Println("Handling connection...")
				go handleConnection(&conn, make(chan *protocol.Command))

				// Sleep for testing purposes
				time.Sleep(time.Second * 2)
				log.Println("Writing data...")

				// Send commands to client for testing purposes
				_, err := protocol.WriteCommands(&conn,
					protocol.NewCommand("SESS", "MP"),
					protocol.NewCommand("LANG", "<>"))
				if err != nil {
					log.Println("Failed to send data:", err)
				}

				// close tcp connection
				log.Println("Closing connection...")
				err = conn.Close()
				if err != nil {
					log.Println("Failed to close connection:", err)
				}

				log.Println("Done!")
			}
		}
	}
}
