package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (server *Server) RegisterRoutes(app *fiber.App) {

	app.Post("/test", func(c *fiber.Ctx) error {
		for key, value := range c.GetReqHeaders() {
			log.Debugf("%s: %s\n", key, value)
		}

		authorization := c.Get("Authorization")
		tempMap := map[string]string{"token": authorization}
		return c.JSON(tempMap)
	})
	//common api
	commonGroup := app.Group("/common", server.api.middleware.ValidatePermission)
	// commonGroup := app.Group("/common")
	commonGroup.Post("/na-list/search", server.SearchListHandle)

}
