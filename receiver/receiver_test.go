package receiver

import "testing"

func TestCreation(t *testing.T) {

	// first case: Invalid Port
	err := StartServer(Host_Information{"Me", "192.168.1.31", "123456", "tcp"})
	err.ToLog().Stdout()
	if err == nil {
		t.Errorf("Invalid Port did not return an error!")
	}

	// second case: Invalid IP
	err = StartServer(Host_Information{"Me", "BakedBeans", "1234", "tcp"})
	err.ToLog().Stdout()
	if err == nil {
		t.Errorf("Invalid IP did not return an error!")
	}

	// third case: Non-numerical port number
	err = StartServer(Host_Information{"Me", "localhost", "Baked Beans", "tcp"})
	err.ToLog().Stdout()
	if err == nil {
		t.Errorf("Non-numerical port number did not return an error!")
	}

	// fourth case: Invalid protocol
	err = StartServer(Host_Information{"Me", "localhost", "1234", "Train"})
	err.ToLog().Stdout()
	if err == nil {
		t.Errorf("Invalid Protocol did not return an error!")
	}

	// test logging to a file
	err.LogToFile()
}
