package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"main/models"
	"main/models/db"
	"strconv"
	"time"
)

func (api *API) Fetch(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"title": api.Config["org.name"],
		"desc":  api.Config["org.desc"],
	})
}

func (api *API) FetchMonitor(c *websocket.Conn) {
	id := c.Params("id")
	idc, _ := strconv.Atoi(id)
	for {
		tickets, err := api.TicketProvider.GetTodayTickets(uint(idc))
		var tickets2 []db.Ticket
		for _, ticket := range tickets {
			if ticket.EndTime.Before(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC)) {
				tickets2 = append(tickets2, ticket)
			}
		}
		if err != nil {
			c.Close()
			return
		}
		c.WriteJSON(tickets2)
		time.Sleep(1 * time.Second)
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

func (api *API) FetchBranches(c *fiber.Ctx) error {
	tickets, err := api.OrgProvider.ListBranches()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(tickets)
}
