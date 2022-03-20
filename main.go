package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/streadway/amqp"
	"encoding/json"
	"log" )

func main() {
	a := app.New()
	w := a.NewWindow("Send RabbitMq")

	hello := widget.NewLabel("Send to RabbitMQ!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Send!", func() {
			log.Printf("hi")
			conn, err := amqp.Dial("amqp://giveout:giveout123@localhost:5672/")
			failOnError(err, "Failed to connect to RabbitMQ")
			defer conn.Close()
		
			ch, err := conn.Channel()
			failOnError(err, "Failed to open a channel")
			defer ch.Close()
		
			err = ch.ExchangeDeclare(
				"giveoutExchange",   // name
				"fanout", // type
				true,     // durable
				false,    // auto-deleted
				false,    // internal
				false,    // no-wait
				nil,      // arguments
			  )

			pass:= &PassModel{
				CardType: "3",
				PassCode:"36096178920526596",
				Serial:"5506",
				Workstation:"wpet-w10-p00320",
			}
			data, _ := json.Marshal(pass)
			body := string(data)
	
			err = ch.Publish(
			"giveoutExchange", // exchange
			"",     // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
			})
			hello.SetText("Sent :)")
		}),
	))

	w.ShowAndRun()
}
func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
  }
  type PassModel struct {
	CardType string
	PassCode string
	Serial string
	Workstation string
  }