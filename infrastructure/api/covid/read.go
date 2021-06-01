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

type ByPrivPub struct {
	Priv *map[string]*ByState
	Pub  *map[string]*ByState
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
		case "BYPRIVPUB":
			filteredDocs := ConvertDataByPubPriv(documents)
			return c.Status(http.StatusOK).JSON(&models.Response{
				Code:  http.StatusOK,
				Error: nil,
				Data:  filteredDocs,
				Count: 1,
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

func ConvertDataByPubPriv(docs *[]*models.Case) *ByPrivPub {
	PubStates := map[string]*ByState{}
	PrivStates := map[string]*ByState{}

	for _, each := range *docs {
		switch each.Origen {
		case "PRIV":
			if state, isOk := PrivStates[each.Entidad_Residencia]; !isOk {
				PrivStates[each.Entidad_Residencia] = &ByState{
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
			break
		case "PUB":
			if state, isOk := PubStates[each.Entidad_Residencia]; !isOk {
				PubStates[each.Entidad_Residencia] = &ByState{
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
			break
		}

	}

	final := ByPrivPub{
		Priv: &PrivStates,
		Pub:  &PubStates,
	}

	return &final
}
