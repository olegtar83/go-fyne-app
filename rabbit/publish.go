package rabbit

import(
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"giveout-emulator/models"
)
func PublishMessage(pass *models.PassModel){
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
}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
  }
