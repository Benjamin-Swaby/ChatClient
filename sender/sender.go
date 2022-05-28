package sender

import (
	"ChatClient/logger"
	"net"
	"time"
)

type Recipient struct {
	Name     string
	Ip       string
	Port     string
	Protocol string
}

func (r Recipient) Send(req string) Sender_error_interface {

	connection, err := net.Dial(r.Protocol, r.Ip+":"+r.Port)
	if err != nil {
		return &Sender_error{4, "Could not establish connection to Recipient", time.Now().Format(time.RFC1123), false}
	}
	///send some data
	_, err = connection.Write([]byte(req))
	logger.Log{
		Text_colour: logger.Green,
		Time:        time.Now().Format(time.RFC1123),
		Msg:         "To:" + connection.RemoteAddr().String() + " -> " + req,
	}.File()

	// confirmation
	buffer := make([]byte, 8192)
	mLen, err := connection.Read(buffer)

	if err != nil {
		return &Sender_error{4, "No confirmation recieved", time.Now().Format(time.RFC1123), false}
	}

	// check confirmation here.

	logger.Log{
		Text_colour: logger.Green,
		Time:        time.Now().Format(time.RFC1123),
		Msg:         "Confirmation Recived from : " + connection.RemoteAddr().String() + " : " + string(buffer[:mLen]),
	}.File()

	defer connection.Close()

	return nil
}
