package infrastructure

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/n3m/covid-dashboard-server/application/services"
	_covidService "github.com/n3m/covid-dashboard-server/application/services/covid"
	"github.com/n3m/covid-dashboard-server/domain/models"
	"github.com/thedevsaddam/gojsonq/v2"
)

type ByState struct {
	State string
	Count int64
	Data  []*models.Case
}

func QueryCustom(c *fiber.Ctx) error {

	filter := _covidService.Filter{}

	err := c.BodyParser(&filter)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(&models.Response{
			Code: http.StatusConflict,
			Error: &models.Error{
				Message: "An unexpected error happened while parsing the params! [00]",
				Error:   err.Error(),
			},
		})
	}

	db, ok := c.Locals("db").(*gojsonq.JSONQ)
	if !ok {
		return c.Status(http.StatusConflict).JSON(&models.Response{
			Code: http.StatusConflict,
			Error: &models.Error{
				Message: "An unexpected error happened! (DB) [01]",
			},
		})
	}

	documents, count, err := services.CovidService.Find(&filter, db)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(&models.Response{
			Code: http.StatusConflict,
			Error: &models.Error{
				Message: "An unexpected error happened while retrieving the data! [02]",
				Error:   err.Error(),
			},
		})
	}

	if filter.ResponseType != nil {
		switch *filter.ResponseType {
		case "BYSTATE":
			filteredDocs := ConvertDataByState(documents)
			return c.Status(http.StatusOK).JSON(&models.Response{
				Code:  http.StatusOK,
				Error: nil,
				Data:  filteredDocs,
				Count: len(*filteredDocs),
			})
		}
	}

	return c.Status(http.StatusOK).JSON(&models.Response{
		Code:  http.StatusOK,
		Error: nil,
		Data:  documents,
		Count: count,
	})
}

func ConvertDataByState(docs *[]*models.Case) *map[string]*ByState {
	states := map[string]*ByState{}

	for _, each := range *docs {
		if state, isOk := states[each.Entidad_Residencia]; !isOk {
			states[each.Entidad_Residencia] = &ByState{
				State: each.Entidad_Residencia,
				Count: 1,
				Data: []*models.Case{
					each,
				},
			}
		} else {
			state.Count++
			state.Data = append(state.Data, each)
		}
	}

	return &states
}
