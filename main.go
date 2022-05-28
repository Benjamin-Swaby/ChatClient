package main

import (
	"ChatClient/receiver"
	"time"
)

// serverInit()
// This procedure will parse a relevant config file and start a listening server
func serverInit() {
	err := receiver.StartServer(receiver.Host_Information{"192.168.1.20", "1234", "tcp"})
	if err != nil {
		err.ToLog().Stdout()
	}
}

// Entry Point
func main() {
	// start the listening server (receiver)
	go serverInit()

	// hang
	for {
		time.Sleep(100 * time.Millisecond)
	}

}
