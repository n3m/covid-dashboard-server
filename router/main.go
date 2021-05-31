package router

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/n3m/covid-dashboard-server/infrastructure"
	"github.com/thedevsaddam/gojsonq"
)

const version = "0.0.1"

//Init ...
func Init(dbFileName string) *fiber.App {
	server := fiber.New(fiber.Config{
		ReadTimeout: time.Second * 15,
	})

	server.Static("/", "./public")

	server.Get("/version", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString(version)
	})

	database := ReadAndParseDatabase(dbFileName)

	CovidEndpoint := infrastructure.NewCovidEndpoint(infrastructure.Input{DB: database})

	server.Get("/covid", CovidEndpoint.FiberHandler())

	return server
}

func ReadAndParseDatabase(fileName string) *gojsonq.JSONQ {
	return gojsonq.New().File(fileName)
}
