package main

import (
	"ChatClient/receiver"
	"ChatClient/sender"
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

	// send a request every 2 seconds to the local server
	for {

		time.Sleep(2 * time.Second)

		sender.Recipient{
			Name:     "Benjamin",
			Ip:       "192.168.1.20",
			Port:     "1234",
			Protocol: "tcp",
		}.Send("Hello World!")
	}

}
