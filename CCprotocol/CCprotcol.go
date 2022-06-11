package CCprotocol

import (
	"ChatClient/sender"
	"strings"
	"time"
)

type CC struct {
	Sender_info sender.Recipient
	Message     string
	Type        string
}

func ParseAsCC(req, protocol string) (CC, CC_error_interface) {
	req_split := strings.Split(req, "/")

	//  req_split[0]			       req_split[1]   req_split[2]
	// FROM:Benjamin:192.168.1.20:1234/Hello World!/Send
	// ^ sender info			       ^message      ^type

	if len(req_split) != 3 {

		//TODO: When users message contains a '\' the request becomes invalid
		// FIX: only use the last element in req_split as the type.
		return CC{sender.Recipient{
				Name:     "",
				Ip:       "",
				Port:     "",
				Protocol: "",
			}, "", ""},
			&CC_parse_error{
				"Failed to Parse request",
				time.Now().Format(time.RFC1123)}
	}

	sender_info_split := strings.Split(req_split[0], ":")

	if len(sender_info_split) != 4 {
		return CC{sender.Recipient{
				Name:     "",
				Ip:       "",
				Port:     "",
				Protocol: "",
			}, "", ""},
			&CC_parse_error{
				"Failed to Parse request - Invalid sender Info",
				time.Now().Format(time.RFC1123)}
	}

	req_sender := sender.Recipient{
		Name:     sender_info_split[1],
		Ip:       sender_info_split[2],
		Port:     sender_info_split[3],
		Protocol: protocol,
	}

	output_cc := CC{req_sender, req_split[1], req_split[2]}

	return output_cc, nil
}

func (req *CC) ToString() string {

	string_resp := "FROM:" +
		req.Sender_info.Name +
		":" + req.Sender_info.Ip +
		":" + req.Sender_info.Port +
		"/" + req.Message +
		"/" + req.Type

	return string_resp
}
