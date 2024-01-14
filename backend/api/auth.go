package api

import (
	"github.com/gofiber/fiber/v2"
	"main/models"
)

func (api *API) AuthLogin(c *fiber.Ctx) error {
	req := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: "Invalid request"})
	}

	acc := api.AccountProvider.New()
	if _, err := acc.Login(req.Username, req.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: err.Error()})
	}
	sess := acc.NewSession()
	return c.JSON(models.ApiLoginResponse{Token: sess})
}

func (api *API) AuthUser(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	return c.JSON(acc.User)
}
