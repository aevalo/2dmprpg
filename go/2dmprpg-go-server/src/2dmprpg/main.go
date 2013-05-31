package main

import (
	"log"
	"2dmprpg/server"
	"os"
	"os/signal"
)

func main() {
	log.Printf("Start application")

	server.Start("0.0.0.0", 8000)

	// catch SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			// intterrupt came handle it
			server.Close()
			log.Printf("Quit application")
		}
	}()

	// main loop
	for {
	}

}
