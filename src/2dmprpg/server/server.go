package server

import (
	"log"
	"net"
	//	"protocol"
	"fmt"
	//	"time"
)

// Package initialization function
func init() {
	// stuff
}

var server *Server

type Server struct {
	Addr  *net.TCPAddr
	Users []*NetUser // TODO: consider a map here
}

type NetUser struct {
	Id            int
	Connection    *net.TCPConn
	Authenticated bool
	Alive         bool
	Channel       chan *Command
}

func (c *NetUser) String() string {
	return fmt.Sprintf("Id: %s, Addr: %s, Authed: %v", c.Id, c.Connection, c.Authenticated)
}

func NewNetUser(conn *net.TCPConn) *NetUser {
	user := new(NetUser)
	user.Alive = true
	user.Authenticated = false
	user.Connection = conn
	user.Channel = make(chan *Command)
	user.Id = len(server.Users) + 1 //change this to proper id
	return user
}

func NewTCPAddr(ip string, port int) *net.TCPAddr {
	addr := new(net.TCPAddr)
	addr.IP = net.ParseIP(ip)
	addr.Port = port
	return addr
}

func HandleConnection(conn *net.TCPConn) {
	//create user
	user := NewNetUser(conn)
	server.Users = append(server.Users, user)

	// handle incoming command the best you can
	for user.Alive {
		cmd := ReadCommand(conn)
		log.Printf("%d: Name: %s, Data: %s\n", user.Id, cmd.Name, cmd.Data)

		// TODO: add some quit / auth handling here

		user.Channel <- cmd
	}

	// close tcp connection
	log.Println("Closing connection", user.Id)
	err := conn.Close()
	if err != nil {
		log.Println("Failed to close connection:", err)
	}

	//TODO: remove from list
}

func Start(ip string, port int) {
	// create server object
	server = new(Server)
	server.Addr = NewTCPAddr(ip, port)

	// start listening tcp connections
	ln, err := net.ListenTCP("tcp", server.Addr)
	if err != nil {
		log.Println("Error occured while listening for connections:", err)
		return
	}
	if ln != nil {
		go func() {
			// waiting for connection
			log.Println("Waiting for incoming connections...")
			conn, err := ln.AcceptTCP()
			if err != nil {
				log.Println("Failed to accept connection:", err)
				return
			} else {
				if conn != nil {
					log.Println("Handling connection... ")
					go HandleConnection(conn)
				}
			}
		}()
		log.Println("Server started")
	}
}

// close all connections 
func Close() {
	for i := 0; i <= len(server.Users); i++ {
		server.Users[i].Alive = false
		server.Users[i].Connection.Close()
		close(server.Users[i].Channel)
	}
	// TODO: add users removal from list
	log.Println("Server closed")
}
