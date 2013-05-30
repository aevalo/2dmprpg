package server

import (
	"log"
	"net"
	"2dmprpg/protocol"
	"fmt"
	//	"time"
)

// Package initialization function
func init() {
	// stuff
}

var server *Server

type Server struct {
	Ip    string
	Port  string
	Users []*NetUser // TODO: consider a map here
}

func (c *Server) HostString() string {
	return fmt.Sprintf("%s:%s", c.Ip, c.Port)
}

type NetUser struct {
	Id            string
	Connection    *net.Conn
	Authenticated bool
	Alive         bool
	Channel       chan *protocol.Command
}

func (c *NetUser) String() string {
	return fmt.Sprintf("Id: %s, Addr: %s, Authed: %v", c.Id, c.Connection.RemoteAddr().String(), c.Authenticated)
}

func NewNetUser(conn *net.Conn) *NetUser {
	user := new(NetUser)
	user.Alive = true
	user.Authenticated = false
	user.Connection = conn
	user.Channel = make(chan *protocol.Command)
	user.Id = conn.RemoteAddr().String() //change this to proper id
	return user
}

func HandleConnection(conn *net.Conn) {
	//create user
	user := NewNetUser(conn)
	append(server.Users, user)

	// handle incoming command the best you can
	for user.Alive {
		cmd := protocol.ReadCommand(conn)
		log.Printf("%s: Name: %s, Data: %s\n", user.Id, cmd.Name, cmd.Data)

		// TODO: add some quit / auth handling here

		user.Channel <- cmd
	}

	// close tcp connection
	log.Println("Closing connection", conn.RemoteAddr().String())
	err = conn.Close()
	if err != nil {
		log.Println("Failed to close connection:", err)
	}
}

func Start(ip, port string) {
	// create server object
	server = new(Server)
	server.Ip = ip
	server.Port = port

	// start listening tcp connections
	ln, err := net.Listen("tcp", server.HostString())
	if err != nil {
		log.Println("Error occured while listening for connections:", err)
		return
	}
	if ln != nil {
		log.Println("Waiting for incoming connections...")
		go func(ln *net.Listener) {
			// waiting for connection
			conn, err := ln.Accept()
			if err != nil {
				log.Println("Failed to accept connection:", err)
				return
			} else {
				if conn != nil {
					log.Println("Handling connection: ", conn.RemoteAddr().String())
					go handleConnection(&conn)
				}
			}
		}(&ln)
		log.Println("Server started")
	}
}

// close all connections 
func Close() {
	for i := 0; i <= len(server.Users); i++ {
		server.Users[i].Alive = false
		server.Users[i].Connection.Close()
		close(server.Users[i].Channel)
		server.Users.Remove(i)
	}
}
