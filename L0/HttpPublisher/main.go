package main

import (
	"encoding/json"
	"fmt"
	"github.com/AnSunn/WBTasks/L0/HttpPublisher/model"
	"github.com/nats-io/stan.go"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	sc                stan.Conn
	infoLogPublisher  = log.New(os.Stdout, "INFO[Publisher] ", log.Ldate|log.Ltime)
	errorLogPublisher = log.New(os.Stderr, "ERROR[Publisher] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func init() {
	// Connect to NATS Streaming server
	var err error
	sc, err = stan.Connect("test-cluster", "http-server-client")
	if err != nil {
		errorLogPublisher.Fatalf("Error connecting to NATS Streaming: %v", err)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// Launch HTTP
	http.HandleFunc("/publish", publishHandler)
	infoLogPublisher.Println("HTTP server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func publishHandler(w http.ResponseWriter, r *http.Request) {
	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var jsonData model.Order
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	//Data validation
	if err := jsonData.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %s", err), http.StatusBadRequest)
		return
	}
	// Publish JSON to NATS Streaming
	subject := "example-channel"
	data, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	err = sc.Publish(subject, data)
	if err != nil {
		http.Error(w, "Error publishing message", http.StatusInternalServerError)
		return
	}

	// Print received JSON
	infoLogPublisher.Printf("Received JSON: %+v\n", jsonData)

	// Send response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message received and published successfully"))
}
