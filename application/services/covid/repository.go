package covid

import (
	"fmt"

	"github.com/n3m/covid-dashboard-server/domain/models"
	"github.com/thedevsaddam/gojsonq/v2"
)

//Find ...
func (c *CovidService) Find(filter *Filter, db *gojsonq.JSONQ) (*[]*models.Case, int64, error) {
	result := []*models.Case{}
	tx := _BuildFilter(filter, db)
	response := tx.Get()

	data, isOk := response.([]interface{})
	if !isOk {
		return nil, 0, fmt.Errorf("Data conversion not ok")
	}

	for _, each := range data {
		c, isOk := each.(map[string]interface{})
		if !isOk {
			return nil, 0, fmt.Errorf("Each Data conversion not ok [Real Type: %T]", each)
		}

		parsedCase := models.Case{}
		if err := parsedCase.Edit(c); err != nil {
			return nil, 0, fmt.Errorf("Couldn't edit case > %w", err)
		}

		result = append(result, &parsedCase)
	}

	return &result, int64(len(result)), nil
}

func _BuildFilter(filter *Filter, db *gojsonq.JSONQ) *gojsonq.JSONQ {
	tx := db

	if filter.Sexo != nil {
		for _, each := range *filter.Sexo {
			if each.Eq != nil {
				tx.Where("sexo", "eq", *each.Eq)
			}

			if each.Ne != nil {
				tx.Where("sexo", "neq", *each.Ne)
			}

			if each.Gte != nil {
				tx.Where("sexo", "gte", *each.Gte)
			}

			if each.Lte != nil {
				tx.Where("sexo", "lte", *each.Lte)
			}
		}
	}

	if filter.Edad != nil {
		for _, each := range *filter.Edad {
			if each.Eq != nil {
				tx.Where("edad", "eq", *each.Eq)
			}

			if each.Ne != nil {
				tx.Where("edad", "neq", *each.Ne)
			}

			if each.Gte != nil {
				tx.Where("edad", "gte", *each.Gte)
			}

			if each.Lte != nil {
				tx.Where("edad", "lte", *each.Lte)
			}
		}
	}

	if filter.Origen != nil {
		for _, each := range *filter.Origen {
			if each.Eq != nil {
				tx.Where("origen", "eq", *each.Eq)
			}

			if each.Ne != nil {
				tx.Where("origen", "neq", *each.Ne)
			}

			if each.Gte != nil {
				tx.Where("origen", "gte", *each.Gte)
			}

			if each.Lte != nil {
				tx.Where("origen", "lte", *each.Lte)
			}
		}
	}

	if filter.Defunto != nil {
		for _, each := range *filter.Defunto {
			if each.Eq != nil {
				tx.Where("fecha_def", "eq", *each.Eq)
			}

			if each.Ne != nil {
				tx.Where("fecha_def", "neq", *each.Ne)
			}

			if each.Gte != nil {
				tx.Where("fecha_def", "gte", *each.Gte)
			}

			if each.Lte != nil {
				tx.Where("fecha_def", "lte", *each.Lte)
			}
		}
	}

	if filter.EstadoResidencia != nil {
		for _, each := range *filter.EstadoResidencia {
			if each.Eq != nil {
				tx.Where("entidad_res", "eq", *each.Eq)
			}

			if each.Ne != nil {
				tx.Where("entidad_res", "neq", *each.Ne)
			}

			if each.Gte != nil {
				tx.Where("entidad_res", "gte", *each.Gte)
			}

			if each.Lte != nil {
				tx.Where("entidad_res", "lte", *each.Lte)
			}
		}
	}

	return tx
}
