package receiver

import (
	"ChatClient/logger"
	"fmt"
	"net"
	"strconv"
	"time"
)

// Information about where to recieve data.
type Host_Information struct {
	Ip      string
	Port    string
	Protcol string
}

// Host_Information.is_valid()
// simple checks of the information to make sure it's worth actually trying to use it
func (hi Host_Information) is_valid() receiver_error_interface {

	port_int, err := strconv.Atoi(hi.Port)
	// is the port numeric
	if err != nil {
		return &receiver_error{
			4,
			"Port is not numeric",
			time.Now().Format(time.RFC1123),
			false}
	}

	// is the port in range
	if port_int > 65535 || port_int < 1 {
		return &receiver_error{
			4,
			"Port is not in valid range",
			time.Now().Format(time.RFC1123),
			false}
	}

	// is the protcol either tcp or udp
	if hi.Protcol != "tcp" && hi.Protcol != "udp" {
		return &receiver_error{
			4,
			"Protocol isn't tcp or udp",
			time.Now().Format(time.RFC1123),
			false}
	}

	return nil
}

// StartServer
// checks to make sure info is valid,
// starts listening for connections on the given port
// handles an incoming connection
func StartServer(info Host_Information) receiver_error_interface {

	valid_error := info.is_valid()

	if valid_error != nil {
		return valid_error
	}

	server, err := net.Listen(info.Protcol, info.Ip+":"+info.Port)

	if err != nil {
		return &receiver_error{
			7,
			"Failed to Start Server",
			time.Now().Format(time.RFC1123),
			true}
	}

	defer server.Close()
	for {
		connection, err := server.Accept()
		if err != nil {
			return &receiver_error{
				8,
				"Couldn't Accept request",
				time.Now().Format(time.RFC1123),
				true}
		}

		// handle connection
		go handleConnection(connection)
	}

	return nil

}

// handleConnection()
// reads the data into a buffer and will send a confirmation back to the client
// also logs any incomming messages and where they were sent from.
func handleConnection(connection net.Conn) {
	// create an input buffer
	buffer := make([]byte, 8192)

	// read io data
	mLen, err := connection.Read(buffer)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	log_message := "From:" + connection.RemoteAddr().String() + " -> " + string(buffer[:mLen])
	logger.Log{logger.Green, time.Now().Format(time.RFC1123), log_message}.Stdout()

	// send Validation back.
	_, err = connection.Write([]byte("Echo From Server : " + connection.RemoteAddr().String() + string(buffer[:mLen])))

	// close the connection
	connection.Close()
}
