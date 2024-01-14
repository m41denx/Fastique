package api

import (
	"github.com/gofiber/fiber/v2"
	"main/models"
	"main/models/db"
)

func (api *API) AdmGetRoles(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	roles, err := api.OrgProvider.ListRoles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(roles)
}

func (api *API) AdmGetUsers(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	users, err := api.OrgProvider.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(users)
}

func (api *API) AdmGetLabels(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	tickets, err := api.OrgProvider.ListLabels()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(tickets)
}

func (api *API) AdmGetBranches(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	tickets, err := api.OrgProvider.ListBranches()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(tickets)
}

func (api *API) AdmCreateBranch(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	branch := &db.Branch{}
	if err := c.BodyParser(branch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: err.Error()})
	}
	if _, err := api.OrgProvider.CreateBranch(branch.Name, map[string]string{}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(branch)
}

func (api *API) AdmCreateSP(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	sp := &struct {
		BranchID uint   `json:"branch"`
		Labels   []uint `json:"labels"`
	}{}
	if err := c.BodyParser(sp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: err.Error()})
	}
	if _, err := api.OrgProvider.CreateSP(sp.BranchID, sp.Labels); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(sp)
}

func (api *API) AdmGetTickets(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	tickets, err := api.TicketProvider.GetTodayTickets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(tickets)
}
