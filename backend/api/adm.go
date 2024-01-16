package api

import (
	"github.com/gofiber/fiber/v2"
	"main/models"
	"main/models/db"
)

// region Roles

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

func (api *API) AdmCreateRole(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	role := &struct {
		Name        string                 `json:"name"`
		Permissions map[string]interface{} `json:"permissions"`
	}{}
	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: err.Error()})
	}
	if _, err := api.OrgProvider.CreateRole(role.Name, role.Permissions, map[string]string{}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(role)
}

// endregion

// region Users

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

func (api *API) AdmCreateUser(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	user := &struct {
		Username string `json:"username"`
		Password string `json:"password"`
		RoleID   uint   `json:"role_id"`
	}{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: err.Error()})
	}
	accx := api.AccountProvider.New()
	if _, err := accx.Register(user.Username, user.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	if err := accx.SetRole(user.RoleID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(user)
}

func (api *API) AdmDeleteUser(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	user, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: err.Error()})
	}
	if err := api.AccountProvider.DeleteUser(uint(user)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return nil
}

func (api *API) AdmUpdateUser(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	user := &struct {
		RoleID uint `json:"role_id"`
	}{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: err.Error()})
	}
	accx := api.AccountProvider.New()
	if err := accx.GetUser(user.RoleID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	if err := accx.SetRole(user.RoleID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(user)
}

// endregion

// region Labels

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

func (api *API) AdmCreateLabel(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	label := &db.Label{}
	if err := c.BodyParser(label); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: err.Error()})
	}
	if _, err := api.OrgProvider.CreateLabel(label.Name, label.Template); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(label)
}

func (api *API) AdmUpdateLabel(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	label := &db.Label{}
	if err := c.BodyParser(label); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: err.Error()})
	}
	if err := api.OrgProvider.UpdateLabel(label); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(label)
}

// endregion

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

// region SPs

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

func (api *API) AdmUpdateSP(c *fiber.Ctx) error {
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
	if err := api.OrgProvider.UpdateSP(sp.BranchID, sp.Labels); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(sp)
}

func (api *API) AdmDeleteSP(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	sp, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{Message: err.Error()})
	}
	if err := api.OrgProvider.DeleteSP(uint(sp)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return nil
}

// endregion

func (api *API) AdmGetTickets(c *fiber.Ctx) error {
	acc := api.AccountProvider.New()
	if authToken(c, acc) != nil || !acc.HasPermission("admin") {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ApiError{Message: "Unauthorized"})
	}
	tickets, err := api.TicketProvider.GetAllTickets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(tickets)
}
