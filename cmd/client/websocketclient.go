package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yurifrl/poc-websocket/pkg/config"
	pb "github.com/yurifrl/poc-websocket/proto"
	"google.golang.org/protobuf/proto"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type WebSocketClient struct {
	config *config.Config
}

func NewWebSocketClient(cfg *config.Config) *WebSocketClient {
	return &WebSocketClient{
		config: cfg,
	}
}

func (client *WebSocketClient) Run() error {
	const maxRetries = 5
	const backoffDuration = 5 * time.Second

	for retries := 0; retries < maxRetries; retries++ {
		if err := client.connectAndRead(); err != nil {
			client.config.Log().WithField("retry", retries+1).WithError(err).Error("Connection failed, retrying...")
			time.Sleep(backoffDuration)
			continue
		}
		break
	}
	return nil
}

func (client *WebSocketClient) connectAndRead() error {
	c, _, err := websocket.DefaultDialer.Dial(client.config.GetEndpoint(), nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	defer c.Close()

	msg := &pb.Message{Content: "Hello, server!"}
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	if err := c.WriteMessage(websocket.BinaryMessage, data); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	for {
		_, response, err := c.ReadMessage()
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		var receivedMsg pb.Message
		if err := proto.Unmarshal(response, &receivedMsg); err != nil {
			return fmt.Errorf("unmarshal: %w", err)
		}
		fmt.Printf("Received from server: %s\n", receivedMsg.GetContent())
	}
}
