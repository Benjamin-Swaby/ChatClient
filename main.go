package main

import (
	"ChatClient/contacts"
	"ChatClient/logger"
	"ChatClient/receiver"
	"time"
)

// serverInit()
// This procedure will parse a relevant config file and start a listening server
func serverInit(HI receiver.Host_Information) {
	logger.Log{logger.Blue, time.Now().Format(time.RFC1123), "Starting Server on:" + HI.Ip + ":" + HI.Port}.File()
	err := receiver.StartServer(HI)
	if err != nil {
		err.ToLog().Stdout()
	}
}

// Entry Point
func main() {
	// start the listening server (receiver)
	go serverInit(receiver.Host_Information{"Benjamin Server", "192.168.1.20", "1234", "tcp"})

	people, C_err := contacts.Import("/home/benjamin/.config/ChatClient/contacts")
	if C_err != nil {
		C_err.LogToFile()
	}

	// send a request every 2 seconds to the local server
	//FROM:NotBenjamin:192.168.1.20:1234/Hello World!/Send

	for {
		time.Sleep(2 * time.Second)
		people[1].Send("FROM:NotBenjamin:192.168.1.20:1234/Hello World!/Send")
	}

}
