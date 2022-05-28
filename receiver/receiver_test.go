package receiver

import "testing"

func TestCreation(t *testing.T) {

	// first case: Invalid Port
	err := StartServer(Host_Information{"localhost", "123456", "tcp"})
	println("Error Recieved : " + err.Error())
	if err == nil {
		t.Errorf("Invalid Port did not return an error!")
	}

	// second case: Invalid IP
	err = StartServer(Host_Information{"Baked Beans", "1234", "tcp"})
	println("Error Recieved : " + err.Error())
	if err == nil {
		t.Errorf("Invalid IP did not return an error!")
	}

	// third case: Non-numerical port number
	err = StartServer(Host_Information{"localhost", "Baked Beans", "tcp"})
	println("Error Recieved : " + err.Error())
	if err == nil {
		t.Errorf("Non-numerical port number did not return an error!")
	}

	// fourth case: Invalid protocol
	err = StartServer(Host_Information{"localhost", "Baked Beans", "Train"})
	println("Error Recieved : " + err.Error())
	if err == nil {
		t.Errorf("Invalid Protocol did not return an error!")
	}

	// test logging to a file
	err.LogToFile()
}
