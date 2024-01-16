package api

import (
	"github.com/gofiber/fiber/v2"
	"main/models"
	"main/models/db"
	"time"
)

func (api *API) WorkerReserve(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}

	var sp db.ServicePoint
	if err := api.OrgProvider.ExposeGorm().First(&sp, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiError{Message: err.Error()})
	}
	if !acc.HasPermission("admin") && float64(sp.BranchID) != acc.User.Role.Privileges["branch"].(float64) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}

	if err := api.OrgProvider.ClaimSP(uint(id), acc.User); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	api.OrgProvider.SetSPAvailability(uint(id), true)

	return c.JSON(sp)
}

func (api *API) WorkerRelease(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	var sp db.ServicePoint
	if err := api.OrgProvider.ExposeGorm().First(&sp, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiError{Message: err.Error()})
	}
	if (!acc.HasPermission("admin") && float64(sp.BranchID) != acc.User.Role.Privileges["branch"].(float64)) || sp.UserID != acc.User.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	if err := api.OrgProvider.UnclaimSP(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}

	api.OrgProvider.SetSPAvailability(uint(id), false)
	return c.JSON(sp)
}

func (api *API) WorkerClaim(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	tid, _ := c.ParamsInt("tid")
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	var sp db.ServicePoint
	if err := api.OrgProvider.ExposeGorm().First(&sp, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiError{Message: err.Error()})
	}
	if (!acc.HasPermission("admin") && float64(sp.BranchID) != acc.User.Role.Privileges["branch"].(float64)) || sp.UserID != acc.User.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}

	t, err := api.TicketProvider.GetTicket(uint(tid))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	t.ServicePointID = sp.ID
	t.ServantID = acc.User.ID
	t.BeginTime = time.Now()
	if _, err := api.TicketProvider.UpdateTicket(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}

	api.OrgProvider.SetSPAvailability(sp.ID, false)

	return c.JSON(sp)
}

func (api *API) WorkerUnclaim(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	tid, _ := c.ParamsInt("tid")
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	var sp db.ServicePoint
	if err := api.OrgProvider.ExposeGorm().First(&sp, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiError{Message: err.Error()})
	}
	if (!acc.HasPermission("admin") && float64(sp.BranchID) != acc.User.Role.Privileges["branch"].(float64)) || sp.UserID != acc.User.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}

	t, err := api.TicketProvider.GetTicket(uint(tid))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	if t.ServicePointID != sp.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	t.EndTime = time.Now()
	if _, err := api.TicketProvider.UpdateTicket(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}

	api.OrgProvider.SetSPAvailability(sp.ID, true)

	return c.JSON(sp)
}

func (api *API) WorkerQueue(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	var sp db.ServicePoint
	if err := api.OrgProvider.ExposeGorm().Preload("Labels").First(&sp, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ApiError{Message: err.Error()})
	}
	if (!acc.HasPermission("admin") && float64(sp.BranchID) != acc.User.Role.Privileges["branch"].(float64)) || sp.UserID != acc.User.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	var tickets []db.Ticket
	for _, lbl := range sp.Labels {
		tts, err := api.TicketProvider.GetTodayTicketsForLabel(sp.BranchID, lbl.ID)
		if err == nil {
			tickets = append(tickets, tts...)
		}
	}

	return c.JSON(tickets)
}
