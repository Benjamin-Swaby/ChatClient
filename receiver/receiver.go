package receiver

import (
	"net"
	"strconv"
	"time"
)

// Information about where to recieve data.
type Host_Information struct {
	ip      string
	port    string
	protcol string
}

func (hi Host_Information) is_valid() receiver_error_interface {

	port_int, err := strconv.Atoi(hi.port)

	if err != nil {
		return &receiver_error{
			4,
			"Port is not numeric",
			time.Now().Format(time.RFC1123),
			false}
	}

	if port_int > 65535 || port_int < 1 {
		return &receiver_error{
			4,
			"Port is not in valid range",
			time.Now().Format(time.RFC1123),
			false}
	}

	if hi.protcol != "tcp" && hi.protcol != "udp" {
		return &receiver_error{
			4,
			"Protocol isn't tcp or udp",
			time.Now().Format(time.RFC1123),
			false}
	}

	return nil
}

func start_server(info Host_Information) receiver_error_interface {

	valid_error := info.is_valid()

	if valid_error != nil {
		return valid_error
	}

	server, err := net.Listen(info.protcol, info.ip+":"+info.port)

	if err != nil {
		return &receiver_error{
			7,
			"Failed to Start Server",
			time.Now().Format(time.RFC1123),
			true}
	}

	defer server.Close()
	for {
		_, err := server.Accept()
		if err != nil {
			return &receiver_error{
				8,
				"Couldn't Accept request",
				time.Now().Format(time.RFC1123),
				true}
		}

		// handle connection
	}

	return nil

}
