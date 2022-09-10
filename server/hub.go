package main

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	Broadcast  chan *Message
	Register   chan *Client
	Unregister chan *Client
	Rooms      map[string]*Room
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *Message, 5),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Rooms:      make(map[string]*Room),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			// Check Room
			if _, isRoomExist := h.Rooms[client.RoomId]; !isRoomExist {
				h.Rooms[client.RoomId] = &Room{
					RoomId:  client.RoomId,
					Clients: make(map[string]*Client),
				}
			}

			room := h.Rooms[client.RoomId]
			if _, isCLientExist := room.Clients[client.ClientId]; !isCLientExist {
				room.Clients[client.ClientId] = client
			}

			log.Println("connection registered")

		case client := <-h.Unregister:
			// Remove the client from the hub
			delete(h.Rooms[client.RoomId].Clients, client.ClientId)
			log.Println("connection unregistered")

		case message := <-h.Broadcast:
			// Message
			if _, exist := h.Rooms[message.RoomId]; exist {
				for _, room := range h.Rooms[message.RoomId].Clients {
					if err := room.Conn.WriteMessage(websocket.TextMessage, []byte(message.Message)); err != nil {
						log.Println("write error:", err)

						room.Conn.WriteMessage(websocket.CloseMessage, []byte{})
						room.Conn.Close()
						delete(h.Rooms[room.RoomId].Clients, message.ClientId)
					}
					if room.RoomId == message.RoomId {
						// Send client
						room.Message <- message
					}
				}
			}

		}
	}
}
