package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"main/providers"
)

type API struct {
	AccountProvider *providers.AccountProvider
	SessionProvider *providers.SessionProvider
	TicketProvider  *providers.TicketProvider

	Host string
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
	app.Get("/antiswagger/*", swagger.HandlerDefault) //Swag

	// region Auth
	auth := app.Group("/auth")
	auth.Post("/login", nil)
	auth.Post("/register", nil)
	//endregion

	return app.Listen(api.Host)
}
