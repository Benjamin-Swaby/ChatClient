package receiver

import (
	"ChatClient/logger"
	"fmt"
	"os"
)

// reveiver interface such that the error can be returned
type receiver_error_interface interface {
	Error() string
	LogToFile()
	Panic_if_panic()
	ToLog() logger.Log
}

type receiver_error struct {
	severity int    //how bad is the error?
	msg      string // what is the error?
	time     string // when was the error?
	panic    bool   // are we going to have to panic over it?
}

// return a string of the time and error message
func (e *receiver_error) Error() string {
	return e.time + " :: " + e.msg
}

// form a log and write it to the log file
// if the error is severe then also print in stdout
func (e *receiver_error) LogToFile() {
	if e.severity < 5 {
		logger.Log{logger.Yellow, e.time, e.msg}.File()
	} else {
		my_log := logger.Log{logger.Red, e.time, e.msg}
		my_log.Stdout()
		my_log.File()
	}
}

// easy panic condition
func (e *receiver_error) Panic_if_panic() {
	if e.panic {
		fmt.Println("Panic! " + e.msg)
		os.Exit(e.severity)
	}
}

// return the error as a log object
func (e *receiver_error) ToLog() logger.Log {
	if e.severity < 5 {
		return logger.Log{logger.Yellow, e.time, e.msg}
	} else {
		return logger.Log{logger.Red, e.time, e.msg}
	}
}
