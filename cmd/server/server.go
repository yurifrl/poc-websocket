package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	// "github.com/gravitational/trace"
	"github.com/gorilla/websocket"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/yurifrl/poc-websocket/pkg/config"
	pb "github.com/yurifrl/poc-websocket/proto"
	"google.golang.org/protobuf/proto"
)

var _ = pp.Println
var upgrader = websocket.Upgrader{}
var log = logrus.New()

type Server struct {
	config *config.Config
}

func main() {
	// Set the default log level to WarnLevel
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stderr)

	cfg, err := config.New(log)
	if err != nil {
		panic(err)
	}
	// Create server
	server := &Server{
		config: cfg,
	}

	//
	http.HandleFunc("/healthz", server.healthz)
	http.HandleFunc("/ws", server.webSocketHandler)

	// Start the HTTP server and log an error if it fails
	log.Info("Starting server on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}

func (s *Server) healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Health, Version:", s.config.GetVersion())
}

func (s *Server) webSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Log the new connection
	log.Info("New WebSocket connection established")

	// Initialize a counter variable
	counter := 0

	// Start a loop to continuously send messages
	for {
		// Increment the counter
		counter++

		// Create a new message with the iterative number
		msg := &pb.Message{Content: fmt.Sprintf("Hello from the server! [%d]", counter)}

		// Marshal the message to protobuf
		data, err := proto.Marshal(msg)
		if err != nil {
			log.Println("marshal:", err)
			break
		}

		// Send the message
		if err := ws.WriteMessage(websocket.BinaryMessage, data); err != nil {
			log.Println("write:", err)
			break
		}

		// Log the sent message
		log.Printf("Sent: %s", msg.GetContent())

		// Wait for a second before sending the next message
		time.Sleep(1 * time.Second)
	}
}
