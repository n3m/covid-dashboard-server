package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thedevsaddam/gojsonq"
)

type CovidEndPoint struct {
	DB *gojsonq.JSONQ
}

type Input struct {
	DB *gojsonq.JSONQ
}

func NewCovidEndpoint(input Input) *CovidEndPoint {
	app := CovidEndPoint(input)
	return &app
}

func (ep *CovidEndPoint) FiberHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		return c.Status(200).JSON("TEST")
	}
}
