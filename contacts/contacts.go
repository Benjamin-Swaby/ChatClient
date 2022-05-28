package contacts

import (
	"ChatClient/sender"
	"fmt"
	"io/ioutil"
)

func Import(dir string) []sender.Recipient, {
	b, err := ioutil.ReadFile(dir)

	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

}
