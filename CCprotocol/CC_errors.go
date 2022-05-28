package CCprotocol

import "ChatClient/logger"

type CC_error_interface interface {
	Error() string
	ToLog() logger.Log
}

type CC_parse_error struct {
	msg  string // what is the error?
	time string // when was the error?
}

func (e *CC_parse_error) Error() string {
	return e.time + " :: " + e.msg
}

func (e *CC_parse_error) ToLog() logger.Log {
	return logger.Log{logger.Yellow_b, e.time, e.msg}
}
