package contacts

import "ChatClient/logger"

type contact_error_interface interface {
	Error() string
	LogToFile()
	ToLog() logger.Log
}

type IO_error struct {
	msg  string
	time string
	dir  string
}

func (e *IO_error) Error() string {
	return e.time + " :: " + e.msg
}

func (e *IO_error) LogToFile() {
	logger.Log{logger.Yellow, e.time, e.msg}.File()
}

func (e *IO_error) ToLog() logger.Log {

}
