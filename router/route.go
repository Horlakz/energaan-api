package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/horlakz/energaan-api/database"
	"github.com/horlakz/energaan-api/handler"
)

func InitializeRouter(router *fiber.App, dbConn database.DatabaseInterface) {

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	// router.Get("*", handler.NotFound)
	router.Get("/", handler.Index)

	main := router.Group("/api/v1")

	main.Get("/monitor", monitor.New(monitor.Config{Title: "Energaan API Monitor"}))

	InitializeAuthRouter(main, dbConn)
	InitializePlanRouter(router, dbConn)
}
