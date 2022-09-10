package main

import (
	"log"
	"time"

	"github.com/gofiber/websocket/v2"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	// The websocket connection.
	Conn *websocket.Conn
	// Buffered channel of outbound messages.
	Message  chan *Message `json:"message"`
	ClientId string        `json:"clientId"`
	RoomId   string        `json:"roomId"`
}

// from webscoket Connections to Hub
func (c *Client) ReadMessage(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}

			return // Calls the deferred function, i.e. closes the connection on error
		}

		msg := Message{
			Message:  string(message),
			ClientId: c.ClientId,
			RoomId:   c.RoomId,
		}
		if messageType == websocket.TextMessage {
			// Broadcast the received message
			h.Broadcast <- &msg
		} else {
			log.Println("websocket message received of type", messageType)
		}
	}
}

// from Hub to websocket Connection
func WriteMessage(c *websocket.Conn) {

}
