package main

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	Broadcast  chan string
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
	Rooms      map[*websocket.Conn]*Room
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan string),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
		Rooms:      make(map[*websocket.Conn]*Room),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Rooms[client] = &Room{}
			log.Println("connection registered")
			log.Print(h.Rooms[client])

		case client := <-h.Unregister:
			// Remove the client from the hub
			delete(h.Rooms, client)
			log.Println("connection unregistered")

		case message := <-h.Broadcast:
			log.Println("message received:", message)
			for room := range h.Rooms {
				if err := room.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("write error:", err)

					room.WriteMessage(websocket.CloseMessage, []byte{})
					room.Close()
					delete(h.Rooms, room)
				}
			}
		}
	}
}
