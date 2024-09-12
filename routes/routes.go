package routes

import (
	"fmt"
	"time"
	_ "time/tzdata"

	"github.com/Thapanut/struct-test/configs"
	_ "github.com/Thapanut/struct-test/docs"
	"github.com/Thapanut/struct-test/routes/controllers"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartServer() {
	t, err := time.LoadLocation(configs.Config.App.TimeZone)
	if err != nil {
		log.Errorf("Error load timezone location. %v\n", err)
		panic(err) // Handle error appropriately
	}
	log.Debug(time.Now())
	time.Local = t
	log.Debug(time.Now())
	app := fiber.New(fiber.Config{
		// BodyLimit: 6 * 1024 * 1024,
	})

	// connect PG
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		ExposeHeaders: "Content-Disposition", // Important for making this header accessible to the client
	}))
	app.Use(logger.New(logger.Config{
		Format:        "[${time}] [${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat:    "02-Jan-2006 15:04:05 MST",
		DisableColors: false,
	}))
	// @title High Common and Sanction List API
	// @version 1.0
	// @description This is a sample swagger for Fiber
	// @termsOfService http://swagger.io/terms/
	// @contact.name API Support
	// @contact.email fiber@swagger.io
	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
	// @BasePath /
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",

		/*
			// Prefill OAuth ClientId on Authorize popup
			OAuth: &swagger.OAuthConfig{
				AppName:  "OAuth Provider",
				ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
			},
			// Ability to change OAuth2 redirect uri location
			OAuth2RedirectUrl: "http://localhost:30003/swagger/oauth2-redirect.html",
		*/
	}))

	// inject trace middleware
	app.Use(otelfiber.Middleware())

	serApp := controllers.NewServer()
	serApp.RegisterRoutes(app)
	// routes.SetupCommissionRouter(app, pg)
	//Show Route
	for _, routes := range app.Stack() {
		for _, route := range routes {
			if route.Method == fiber.MethodGet || route.Method == fiber.MethodPost {
				fmt.Println(route.Method + ":" + route.Path)
			}
		}
	}

	if err := app.Listen(fmt.Sprintf(":%d", configs.Config.App.Port)); err != nil {
		panic(fmt.Errorf("unable to start fiber: %s", err.Error()))
	}

}
