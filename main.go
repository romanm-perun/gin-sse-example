// Example SSE server in Golang.
//     $ go run main.go

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	br "github.com/romanm-perun/gin-sse-example/broker"
)

func main() {
	broker := br.NewBroker()

	router := gin.Default()
	router.GET("/subscription/:topic", broker.ServeHTTP)

	// Set it running - listening and broadcasting events
	go broker.Listen()

	// Emitting topic A events to broker
	go func(topic string) {
		for {
			time.Sleep(time.Second * 2)
			eventString := fmt.Sprintf("the time is %v", time.Now())
			log.Println("Emitting event for " + topic)
			broker.Notifier <- br.NotificationEvent{
				EventName: topic,
				Payload:   eventString,
			}
		}
	}("topic A")

	// Emitting topic B events to broker
	go func(topic string) {
		for {
			time.Sleep(time.Millisecond * 500)
			eventString := fmt.Sprintf("the UTC time is %v", time.Now().UTC())
			log.Println("Emitting event for " + topic)
			broker.Notifier <- br.NotificationEvent{
				EventName: topic,
				Payload:   eventString,
			}
		}
	}("topic B")

	log.Fatal("HTTP server error: ", router.Run(":3000"))
}
