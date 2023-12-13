package stan

import (
	"encoding/json"
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/cache"
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/model"
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/storage"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"sync"
)

var (
	infoLogStan  = log.New(os.Stdout, "INFO[STAN] ", log.Ldate|log.Ltime)
	errorLogStan = log.New(os.Stderr, "ERROR[STAN] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogData = log.New(os.Stderr, "ERROR[DATA}\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogData  = log.New(os.Stdout, "INFO[DATA] ", log.Ldate|log.Ltime)
)

// Stan function
func NATSStreaming(wg *sync.WaitGroup) {
	defer wg.Done()
	// Connect to NATS Streaming server
	sc, err := stan.Connect("test-cluster", "subscriber-client")
	if err != nil {
		errorLogStan.Fatal(err)
	}
	defer sc.Close()

	// Subscribe to "example-channel" adding durable pack to pick up where we left off from an earlier session
	//Doing this causes the NATS Streaming server to track the last acknowledged message for that clientID + durable name,
	//so that only messages since the last acknowledged message will be delivered to the client.
	subject := "example-channel"
	_, err = sc.Subscribe(subject, handleMessage, stan.DurableName("my-durable"))
	if err != nil {
		errorLogStan.Fatal(err)
	}

	infoLogStan.Println("Subscriber is waiting for messages. Press Ctrl+C to exit.")
	//Subscriber always waits messages
	<-make(chan struct{})
}

// Handle Messages func
func handleMessage(msg *stan.Msg) {
	var order model.Order
	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		errorLogData.Println("Error decoding JSON:", err)
		return
	}

	// Validate order
	if !order.Validate() {
		errorLogData.Printf("Invalid order, please correct sent JSON. You've sent following: %+v", order)
		return
	}

	//Add order to db
	if err := storage.InsertOrder(order); err != nil {
		storage.ErrorLogDB.Printf("Unable to insert the order to db: %v", err)
		return
	}

	//Update cache to be able to receive data about just sent order by order_id
	cache.UpdateCache(order)
	infoLogData.Printf("Received valid order: %+v\n", order)
}
