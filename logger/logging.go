package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	Red      = "\u001b[31m"
	Green    = "\u001b[32m"
	Blue     = "\u001b[34m"
	Yellow   = "\u001b[33m"
	Red_b    = "\u001b[41m"
	Green_b  = "\u001b[42m"
	Blue_b   = "\u001b[44m"
	Yellow_b = "\u001b[43m"
	Reset    = "\u001b[0m"
)

type Log struct {
	Text_colour string
	Time        string
	Msg         string
}

func (l Log) Stdout() {
	out := Blue + l.Time + Reset + " :: " + l.Text_colour + l.Msg + Reset
	fmt.Println(out)
}

func (l Log) File() {
	out := Blue + l.Time + Reset + " :: " + l.Text_colour + l.Msg + Reset + "\n"

	f, err := os.OpenFile(".ChatClient/Chatlogs.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(out); err != nil {
		log.Println(err)
	}
}
