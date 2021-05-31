package router

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

const version = "0.0.1"

//Init ...
func Init() *fiber.App {

	server := fiber.New(fiber.Config{
		ReadTimeout: time.Second * 15,
	})

	return server
}

func SetRoutes(server *fiber.App) {
	/* STATIC */
	server.Static("/", "./public")

	/* INFO */
	server.Get("/version", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString(version)
	})

	/* CUSTOM ENDPOINTS */
	// server.Post("/covid/*", covid.QueryCustom)
}
