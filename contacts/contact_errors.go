package contacts

import "ChatClient/logger"

type Contact_error_interface interface {
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
	logger.Log{logger.Yellow, e.time, e.msg + "(" + e.dir + ")"}.File()
}

func (e *IO_error) ToLog() logger.Log {
	return logger.Log{logger.Yellow, e.time, e.msg + "(" + e.dir + ")"}
}
