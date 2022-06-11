package sender

import (
	"ChatClient/logger"
	"crypto/sha256"
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
		return &Sender_error{4, "No confirmation received", time.Now().Format(time.RFC1123), false}
	}

	original_hash := sha256.Sum256([]byte(req))
	original_hash_s := string(original_hash[:])
	received_hash := string(buffer[:mLen])

	if original_hash_s == received_hash {
		logger.Log{
			Text_colour: logger.Green,
			Time:        time.Now().Format(time.RFC1123),
			Msg:         "Confirmation Received from : " + connection.RemoteAddr().String(),
		}.File()
	} else {
		logger.Log{
			Text_colour: logger.Yellow_b,
			Time:        time.Now().Format(time.RFC1123),
			Msg:         "Invalid Confirmation Received from: " + connection.RemoteAddr().String(),
		}.File()
	}

	defer connection.Close()

	return nil
}
