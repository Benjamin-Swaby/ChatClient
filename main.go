package main

import (
	"ChatClient/CCprotocol"
	"ChatClient/contacts"
	"ChatClient/logger"
	"ChatClient/receiver"
	"ChatClient/sender"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strings"
	"time"
)

// serverInit()
// This procedure will parse a relevant config file and start a listening server
func serverInit(HI receiver.Host_Information, CE *fyne.Container) {
	logger.Log{logger.Blue, time.Now().Format(time.RFC1123), "Starting Server on:" + HI.Ip + ":" + HI.Port}.File()
	err := receiver.StartServer(HI, CE)
	if err != nil {
		err.ToLog().Stdout()
	}
}

// Entry Point
func main() {

	conf := receiver.Host_Information{"Benjamin", "192.168.1.20", "1234", "tcp"}

	people, C_err := contacts.Import("/home/benjamin/.config/ChatClient/contacts")
	if C_err != nil {
		C_err.LogToFile()
	}

	target := people[0]

	// send a request every 2 seconds to the local server
	//FROM:NotBenjamin:192.168.1.20:1234/Hello World!/Send

	a := app.New()
	w := a.NewWindow("Chat Client Version:no")

	w.SetContent(widget.NewLabel("Chat Client"))

	// ChatFeild
	chat_elems := container.NewVBox()
	Chat := container.NewScroll(chat_elems)

	// input_bar
	input := widget.NewEntry()
	input.PlaceHolder = "Message : " + target.Name
	send := widget.NewButton("Send ->", func() {

		if input.Text[0] == '/' {
			msg := strings.Split(input.Text, " ")
			switch msg[0] {
			case "/target":
				// switch the target and display at the top
				// load messages from that client from the text file
				// display messages in chat_elems

				if len(msg) != 2 {
					chat_elems.Add(canvas.NewText("Usage : /target [target name]", color.RGBA{255, 0, 0, 0}))
				} else {

					found := false
					for _, person := range people {
						if person.Name == msg[1] {
							target = person
							found = true
						}
					}

					if !found {
						chat_elems.Add(canvas.NewText("Not a Target", color.RGBA{255, 0, 0, 0}))
					} else {

						// set chat here.
						chat_elems.Add(canvas.NewText("--- "+target.Name+" ---", color.RGBA{255, 255, 0, 0}))
						input.PlaceHolder = "Message : " + target.Name
					}

				}

			case "/me":
				chat_elems.Add(canvas.NewText("Name: "+conf.Name, color.RGBA{0, 0, 255, 0}))
				chat_elems.Add(canvas.NewText("Ip : "+conf.Ip, color.RGBA{0, 0, 255, 0}))
				chat_elems.Add(canvas.NewText("Port :"+conf.Port, color.RGBA{0, 0, 255, 0}))
				chat_elems.Add(canvas.NewText("Protcol :"+conf.Protcol, color.RGBA{0, 0, 255, 0}))

			case "/list":
				for _, person := range people {

					strout := person.Name + " on " + person.Ip + ":" + person.Port

					chat_elems.Add(canvas.NewText(strout, color.RGBA{255, 0, 255, 0}))
				}
			}

		} else {

			// if not a client command - send the message

			chat_elems.Add(canvas.NewText(time.Now().Format(time.RFC1123)+" : "+conf.Name+" : "+input.Text, color.White))

			req := CCprotocol.CC{
				Sender_info: sender.Recipient{conf.Name, conf.Ip, conf.Port, conf.Protcol},
				Message:     input.Text,
				Type:        "Send",
			}

			target.Send(req.ToString())

		}

	})

	input_bar := container.New(layout.NewBorderLayout(nil, nil, nil, send), send, input)
	// final render layout
	content := container.New(layout.NewBorderLayout(nil, input_bar, nil, nil), input_bar, Chat)

	// start the listening server (receiver)
	go serverInit(conf, chat_elems)

	w.SetContent(content)
	w.Resize(fyne.NewSize(500, 300))
	w.ShowAndRun()

}
