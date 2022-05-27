package receiver

import (
	"log"
	"os"
)

type receiver_error_interface interface {
	Error() string
	LogToFile()
}

type receiver_error struct {
	severity int
	msg      string
	time     string
	panic    bool
}

func (e *receiver_error) Error() string {
	return e.time + " :: " + e.msg
}

func (e *receiver_error) LogToFile() {

	output_string := e.time + " :: " + e.msg + "\n"

	f, err := os.OpenFile("receiver.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(output_string); err != nil {
		log.Println(err)
	}

}
