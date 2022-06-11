package receiver

import (
	"ChatClient/CCprotocol"
	"ChatClient/logger"
	"crypto/sha256"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"log"
	"net"
	"os"
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
func StartServer(info Host_Information, CE *fyne.Container) receiver_error_interface {

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
		go handleConnection(connection, info.Protcol, info, CE)
	}

	return nil

}

func Write_message(msg string, cc_obj CCprotocol.CC) {

	f, err := os.OpenFile(".ChatClient/msgs/"+cc_obj.Sender_info.Name+"."+cc_obj.Sender_info.Ip,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(msg); err != nil {
		log.Println(err)
	}

}

// handleConnection()
// reads the data into a buffer,
// parses the request,
// sends back a response
func handleConnection(connection net.Conn, protocol string, h_i Host_Information, CE *fyne.Container) {
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

	//TODO do something with cc_obj i.e write to a UI or anything!
	println(cc_obj.Sender_info.Name + " : " + cc_obj.Message)

	// send notification
	go func() {
		CE.Add(canvas.NewText(time.Now().Format(time.RFC1123)+" : "+cc_obj.Sender_info.Name+" : "+cc_obj.Message, color.RGBA{255, 165, 0, 0}))
		time.Sleep(2 * time.Second)
		CE.RemoveAll()
	}()

	go Write_message(time.Now().Format(time.RFC1123)+" : "+cc_obj.Sender_info.Name+" : "+cc_obj.Message+"\n", cc_obj)

	// send Validation back.
	hash := sha256.Sum256([]byte(req))
	_, err = connection.Write([]byte(string(hash[:])))

	connection.Close()
}
