package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (api *API) Fetch(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"title": api.Config["org.name"],
		"desc":  api.Config["org.desc"],
	})
}

func (api *API) FetchMonitor(c *websocket.Conn) {
	id := c.Params("id")
	for {
		c.WriteMessage(websocket.TextMessage, []byte(id))
	}
}
