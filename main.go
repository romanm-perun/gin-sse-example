package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/romanm-perun/gin-sse-example/broker"
)

// Example SSE server in Golang.
//     $ go run main.go

func main() {
	broker := broker.NewBroker()

	router := gin.Default()
	router.GET("/", broker.ServeHTTP)

	// Set it running - listening and broadcasting events
	go broker.Listen()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			eventString := fmt.Sprintf("the time is %v", time.Now())
			log.Println("Receiving event")
			broker.Notifier <- []byte(eventString)
		}
	}()

	log.Fatal("HTTP server error: ", router.Run(":3000"))
}
