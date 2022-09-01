package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

// var clients = make(map[*websocket.Conn]Client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
// var register = make(chan *websocket.Conn)
// var broadcast = make(chan string)
// var unregister = make(chan *websocket.Conn)

func main() {
	app := fiber.New()
	hub := NewHub()
	// middleware
	app.Use(cors.New())
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.Contains(c.Route().Path, "/ws")
		},
	}))

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	go hub.Run()

	app.Get("/ws/global", websocket.New(func(c *websocket.Conn) {
		hub.Register <- c
		ReadMessage(c, hub)
	}))
	log.Fatal(app.Listen(":3000"))
}
