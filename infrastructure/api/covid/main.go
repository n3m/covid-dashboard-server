package infrastructure

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/n3m/covid-dashboard-server/domain/models"
)

func QueryCustom(c *fiber.Ctx) error {

	params := models.FilterArguments{}

	err := c.BodyParser(&params)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(&models.Response{
			Code: http.StatusConflict,
			Error: &models.Error{
				Message: "An unexpected error happened while parsing the params! [00]",
				Error:   err.Error(),
			},
		})
	}

	return c.Status(200).JSON("TEST")
}
