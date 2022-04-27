package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/QuickSocket/qs-forward/model"
	"github.com/gorilla/websocket"
)

type WebSocket struct {
	clientId     string
	clientSecret string
	webSocketURL string
	callbackc    chan<- *model.Callback
}

func NewWebSocket(clientId, clientSecret, webSocketURL string, callbackc chan<- *model.Callback) *WebSocket {
	return &WebSocket{
		clientId:     clientId,
		clientSecret: clientSecret,
		webSocketURL: webSocketURL,
		callbackc:    callbackc,
	}
}

func (s *WebSocket) Start(logger *log.Logger) error {
	for {
		logger.Printf("Attempting connection to environment...\n")

		conn, err := attemptConnection(s.clientId, s.clientSecret, s.webSocketURL)
		if err != nil {
			return err
		}

		logger.Printf("Connected!\n")

		for {
			if err := receiveAndHandleMessage(conn, s.callbackc); err != nil {
				logger.Printf("Closing connection to environment due to: %v\n", err)
				break
			}
		}

		conn.Close()
	}
}

func attemptConnection(clientId, clientSecret, webSocketURL string) (*websocket.Conn, error) {
	headers := http.Header{}
	headers.Add("Authorization", produceAuthHeader(clientId, clientSecret))

	conn, res, err := websocket.DefaultDialer.Dial(webSocketURL, headers)
	if err != nil {
		return nil, fmt.Errorf("failed to connect WebSocket (%v): %w", res.Status, err)
	}

	return conn, nil
}

func produceAuthHeader(clientId, clientSecret string) string {
	value := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", clientId, clientSecret)))
	return fmt.Sprintf("Basic %v", value)
}

func receiveAndHandleMessage(conn *websocket.Conn, callbackc chan<- *model.Callback) error {
	msgType, msg, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	if msgType != websocket.TextMessage {
		return fmt.Errorf("expected a message type of %v but was %v", websocket.TextMessage, msgType)
	}

	callbackModel := &model.Callback{}
	if err := json.Unmarshal(msg, callbackModel); err != nil {
		return fmt.Errorf("failed to unmarshal callback payload: %w", err)
	}

	callbackc <- callbackModel
	return nil
}
