package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"main/models"
	"main/models/db"
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

func (api *API) FetchLabels(c *fiber.Ctx) error {
	tickets, err := api.OrgProvider.ListLabels()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(tickets)
}

func (api *API) FetchBranchesWithLabel(c *fiber.Ctx) error {
	lbl, _ := c.ParamsInt("label")
	branches, err := api.OrgProvider.ListBranches()
	var branches2 []db.Branch
	for _, branch := range branches {
		found := false
		for _, sp := range branch.ServicePoints {
			for _, l := range sp.Labels {
				if l.ID == uint(lbl) {
					branches2 = append(branches2, branch)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(branches2)
}
