package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"main/providers"
)

type API struct {
	AccountProvider *providers.AccountProvider
	TicketProvider  *providers.TicketProvider
	OrgProvider     *providers.OrgProvider

	Config map[string]string
	Host   string
}

func StartWS(api *API) error {
	app := fiber.New(fiber.Config{
		BodyLimit:     5 * 1024 * 1024, // 5MB Body Limit?
		CaseSensitive: true,
		ServerHeader:  "Fiber",
	})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	app.Use(recover.New())

	// region Monitor
	app.Use("/monitor", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/monitor/:id", websocket.New(nil))
	//endregion

	// region Auth
	auth := app.Group("/auth")
	auth.Post("/login", api.AuthLogin)
	auth.Get("/user", api.AuthUser)
	//endregion

	//region Fetch
	fetch := app.Group("/fetch")
	fetch.Get("/org", api.Fetch)
	fetch.Get("/labels", api.FetchLabels)
	fetch.Get("/branches/:label", api.FetchBranchesWithLabel)
	//endregion

	//region Adm
	adm := app.Group("/adm")
	adm.Get("/roles", api.AdmGetRoles)

	adm.Get("/users", api.AdmGetUsers)

	adm.Post("/user", api.AdmCreateUser)
	adm.Post("/user/:id", api.AdmUpdateUser)
	adm.Delete("/user/:id", api.AdmDeleteUser)

	adm.Get("/branches", api.AdmGetBranches) // Get Branch tree
	adm.Post("/branch", api.AdmCreateBranch) // Create branch
	//adm.Post("/branch/:id", nil) // Edit branch

	adm.Post("/sp", api.AdmCreateSP)
	adm.Post("/sp/:id", api.AdmUpdateSP)
	adm.Delete("/sp/:id", api.AdmDeleteSP)

	adm.Get("/labels", api.AdmGetLabels)
	adm.Post("/label", api.AdmCreateLabel)
	adm.Post("/label/:id", api.AdmUpdateLabel)

	adm.Get("/tickets", api.AdmGetTickets)
	//endregion

	return app.Listen(api.Host)
}

func authToken(c *fiber.Ctx, acc *providers.Account) error {
	token := c.Get("Authorization")
	err := acc.AuthSession(token)
	return err
}
