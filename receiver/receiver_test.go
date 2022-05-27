package receiver

import "testing"

func TestCreation(t *testing.T) {

	// first case: Invalid Port
	err := start_server(Host_Information{"localhost", "123456", "tcp"})
	println("Error Recieved : " + err.Error())
	if err == nil {
		t.Errorf("Invalid Host Information did not return an error!")
	}

	// second case: Invalid IP
	err = start_server(Host_Information{"Baked Beans", "1234", "tcp"})
	println("Error Recieved : " + err.Error())
	if err == nil {
		t.Errorf("Invalid Host Information did not return an error!")
	}

	// third case: Non-numerical port number
	err = start_server(Host_Information{"localhost", "Baked Beans", "tcp"})
	println("Error Recieved : " + err.Error())
	if err == nil {
		t.Errorf("Invalid Host Information did not return an error!")
	}

	// test logging to a file
	err.LogToFile()
}
