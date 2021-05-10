package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/romanm-perun/gin-sse-example/broker"
)

// Example SSE server in Golang.
//     $ go run sse.go

func main() {
	broker := broker.NewServer()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			eventString := fmt.Sprintf("the time is %v", time.Now())
			log.Println("Receiving event")
			broker.Notifier <- []byte(eventString)
		}
	}()

	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:3000", broker))
}
