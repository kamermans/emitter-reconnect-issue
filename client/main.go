package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	emitter "github.com/emitter-io/go/v2"
)

const (
	broker     = "ws://emitter-test-server:8080"
	channel    = "recontest/chan"
	channelKey = "jBDJnSkZqaMeRVzGGzHU1-xMGPXFcD2j"
)

var (
	emitterClient *emitter.Client
	action        = ""
)

func init() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <send|receive>\n", os.Args[0])
		os.Exit(2)
	}

	action = os.Args[1]
	if action != "send" && action != "receive" {
		fmt.Printf("Usage: %s <send|receive>\n", os.Args[0])
		os.Exit(3)
	}

	mqtt.DEBUG = log.New(os.Stderr, "", 0)
	mqtt.ERROR = log.New(os.Stderr, "", 0)
}

func main() {
	fmt.Println("Delaying client startup for 5 seconds...")
	time.Sleep(5 * time.Second)

	// Create the emitter client
	emitterClient = emitter.NewClient(
		emitter.WithAutoReconnect(true),
		emitter.WithConnectTimeout(10*time.Second),
		emitter.WithPingTimeout(30*time.Second),
		emitter.WithKeepAlive(60*time.Second),
		emitter.WithMaxReconnectInterval(120*time.Second),
		emitter.WithBrokers(broker),
	)

	if connErr := emitterClient.Connect(); connErr != nil {
		log.Fatalf("Error on Client.Connect(): %v", connErr)
	}

	switch action {
	case "send":
		doSend()
	case "receive":
		doReceive()
	}

	emitterClient.Disconnect(2 * time.Second)
}

func doSend() {
	i := 0
	for {
		// Send message to the channel
		message := fmt.Sprintf("test msg #%v", i)
		log.Printf("Sending message '%v'", message)
		emitterClient.Publish(channelKey, channel, message)
		time.Sleep(10 * time.Second)
		i++
	}
}

func doReceive() {

	// Subscribe to the channel(s)
	emitterClient.Subscribe(channelKey, channel, func(client *emitter.Client, msg emitter.Message) {
		log.Printf("Message received: [%s] %s", msg.Topic(), msg.Payload())
	})

	// Wait indefinitely
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
