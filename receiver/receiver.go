package receiver

import (
	"ChatClient/CCprotocol"
	"ChatClient/logger"
	"net"
	"strconv"
	"time"
)

// Information about where to recieve data.
type Host_Information struct {
	Name    string
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
			"Failed to Start Server: " + err.Error(),
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
		go handleConnection(connection, info.Protcol)
	}

	return nil

}

// handleConnection()
// reads the data into a buffer,
// parses the request,
// sends back a response
func handleConnection(connection net.Conn, protocol string) {
	// create an input buffer
	buffer := make([]byte, 8192)

	// read io data
	mLen, err := connection.Read(buffer)

	if err != nil {
		logger.Log{logger.Yellow_b, time.Now().Format(time.RFC1123), "Failed to read data from Connection"}.Stdout()
	}

	// log incoming message
	log_message := "From:" + connection.RemoteAddr().String() + " -> " + string(buffer[:mLen])
	logger.Log{logger.Green, time.Now().Format(time.RFC1123), log_message}.File()

	// parse Request
	req := string(buffer[:mLen])

	cc_obj, CCerr := CCprotocol.ParseAsCC(req, protocol)
	if CCerr != nil {
		CCerr.ToLog().File()
	}

	// do something with cc_obj
	println(cc_obj.Sender_info.Name + " : " + cc_obj.Message)

	// send Validation back.
	_, err = connection.Write([]byte("Received!"))

	// close the connection
	connection.Close()
}
