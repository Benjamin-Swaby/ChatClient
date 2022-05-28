package contacts

import (
	"ChatClient/logger"
	"ChatClient/sender"
	"io/ioutil"
	"strings"
	"time"
)

// Import()
// This function will import contancts from a directory
func Import(dir string) ([]sender.Recipient, Contact_error_interface) {
	b, err := ioutil.ReadFile(dir)

	if err != nil {
		return nil, &IO_error{
			msg:  "Failed to Read File",
			time: time.Now().Format(time.RFC1123),
			dir:  dir,
		}
	}

	file_contents := string(b)
	entries := strings.Split(file_contents, "\n")

	Recipients := make([]sender.Recipient, len(entries))

	for i, entry := range entries {
		if len(entry) < 4 {
			continue
		}

		feilds := strings.Split(entry, " ")

		if len(feilds) != 4 {
			logger.Log{
				Text_colour: logger.Red,
				Time:        time.Now().Format(time.RFC1123),
				Msg:         "Invalid Contact Entry in: " + dir,
			}.File()
		}

		Recipients[i] = sender.Recipient{feilds[0], feilds[1], feilds[2], feilds[3]}
	}

	return Recipients, nil
}
