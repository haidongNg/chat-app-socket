package main

import (
	"bytes"
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

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadMessage(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}

			return // Calls the deferred function, i.e. closes the connection on error
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

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

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WriteMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		log.Println("Connection was closed")
	}()
	for {
		select {
		case message, ok := <-c.Message:
			if !ok {
				return
			}

			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

			if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(message.Message)); err != nil {
				log.Println("write error:", err)
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				c.Conn.Close()
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, c *websocket.Conn) {
	client := &Client{
		Conn:     c,
		RoomId:   c.Query("roomId"),
		ClientId: c.Query("clientId"),
		Message:  make(chan *Message, 5),
	}
	// Register
	hub.Register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WriteMessage()
	client.ReadMessage(hub)
}
