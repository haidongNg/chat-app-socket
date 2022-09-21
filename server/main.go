package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

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
		serveWs(hub, c)
	}))
	log.Fatal(app.Listen(":3000"))
}
